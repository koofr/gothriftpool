package gothriftpooltest

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/koofr/go-resourcepool"
	"github.com/koofr/thriftutils"

	"github.com/koofr/gothriftpool"
	"github.com/koofr/gothriftpool/gothriftpooltest/myservice"
	"github.com/koofr/gothriftpool/gothriftpooltest/myserviceproxy"
)

func TestGenerator(t *testing.T) {
	iface := "myservice.MyService"

	generator, err := gothriftpool.NewGenerator(iface)
	if err != nil {
		t.Fatal(err)
	}

	code, err := generator.Generate()
	if err != nil {
		t.Fatal(err)
	}

	testGeneratorExpected, err := ioutil.ReadFile("myserviceproxy/myserviceproxy.go")
	if err != nil {
		t.Fatal(err)
	}

	if string(code) != string(testGeneratorExpected) {
		t.Fatal("Generated code doesn't not match expected (./myserviceproxy/myserviceproxy.go):\n\n" + string(code))
	}
}

func TestProxy(t *testing.T) {
	client, serverClose, poolClose := initProxy(t, defaultPing, defaultGetResponse)
	defer serverClose()
	defer poolClose()

	res, err := client.GetResponse(context.Background(), "id", &myservice.MyRequest{
		Req: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Res != "test" {
		t.Fatalf("expected response to be test, got %s", res.Res)
	}
}

func TestProxyContextTimeout(t *testing.T) {
	i := 0
	client, serverClose, poolClose := initProxy(t, defaultPing, func(ctx context.Context, id myservice.UUID, req *myservice.MyRequest) (r *myservice.MyResponse, err error) {
		if i == 0 {
			time.Sleep(500 * time.Millisecond)
		}
		i++
		return defaultGetResponse(ctx, id, req)
	})
	defer serverClose()
	defer poolClose()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// timeout is ignored in thrift client but cancelled context can produce a
	// broken client
	res, err := client.GetResponse(timeoutCtx, "id", &myservice.MyRequest{
		Req: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Res != "test" {
		t.Fatalf("expected response to be test, got %s", res.Res)
	}

	res, err = client.GetResponse(context.Background(), "id", &myservice.MyRequest{
		Req: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Res != "test" {
		t.Fatalf("expected response to be test, got %s", res.Res)
	}
}

func TestProxyContextCancel(t *testing.T) {
	client, serverClose, poolClose := initProxy(t, defaultPing, defaultGetResponse)
	defer serverClose()
	defer poolClose()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := client.GetResponse(ctx, "id", &myservice.MyRequest{
		Req: "test",
	})
	if err != context.Canceled {
		t.Fatalf("expected error to be context.Canceled but got %v", err)
	}
}

func TestProxyPoolClosed(t *testing.T) {
	client, serverClose, poolClose := initProxy(t, defaultPing, defaultGetResponse)
	defer serverClose()
	poolClose()

	_, err := client.GetResponse(context.Background(), "id", &myservice.MyRequest{
		Req: "test",
	})
	var errCreateResource *myserviceproxy.ErrCreateResource
	if !errors.As(err, &errCreateResource) {
		t.Fatalf("expected error to be ErrCreateResource but got %v", err)
	}
	if !errors.Is(err, resourcepool.ErrPoolClosed) {
		t.Fatalf("expected error to be resourcepool.ErrPoolClosed but got %v", err)
	}
}

func TestProxyServerClosed(t *testing.T) {
	client, serverClose, poolClose := initProxy(t, defaultPing, defaultGetResponse)
	serverClose()
	defer poolClose()

	_, err := client.GetResponse(context.Background(), "id", &myservice.MyRequest{
		Req: "test",
	})
	var errCreateResource *myserviceproxy.ErrCreateResource
	if !errors.As(err, &errCreateResource) {
		t.Fatalf("expected error to be ErrCreateResource but got %v", err)
	}
	if !strings.Contains(err.Error(), "dial tcp") {
		t.Fatalf("expected error to contain dial tcp but got %v", err)
	}
}

func TestProxyPingError(t *testing.T) {
	client, serverClose, poolClose := initProxy(t, func(ctx context.Context) (err error) {
		return fmt.Errorf("custom ping error")
	}, defaultGetResponse)
	defer serverClose()
	defer poolClose()

	_, err := client.GetResponse(context.Background(), "id", &myservice.MyRequest{
		Req: "test",
	})
	var errCreateResource *myserviceproxy.ErrCreateResource
	if !errors.As(err, &errCreateResource) {
		t.Fatalf("expected error to be ErrCreateResource but got %v", err)
	}
	if !strings.Contains(err.Error(), "custom ping error") {
		t.Fatalf("expected error to contain custom ping error but got %v", err)
	}
}

func initProxy(
	t *testing.T,
	ping func(ctx context.Context) (err error),
	getResponse func(ctx context.Context, id myservice.UUID, req *myservice.MyRequest) (r *myservice.MyResponse, err error),
) (myservice.MyService, func(), func()) {
	host := "127.0.0.1"
	l, err := net.Listen("tcp", host+":0")
	if err != nil {
		t.Fatal(err)
	}
	port := l.Addr().(*net.TCPAddr).Port
	l.Close()

	addr := net.JoinHostPort(host, strconv.Itoa(port))

	serverClose, err := createServer(addr, &handler{
		ping:        ping,
		getResponse: getResponse,
	})
	if err != nil {
		t.Fatal(err)
	}

	pool := myserviceproxy.NewClientPool(
		func(ctx context.Context) (myservice.MyService, func(), error) {
			return createClient(addr)
		},
		10,
		10,
	)
	client := myserviceproxy.New(pool.GetClient)

	return client, serverClose, pool.Close
}

func defaultPing(ctx context.Context) (err error) {
	return nil
}

func defaultGetResponse(ctx context.Context, id myservice.UUID, req *myservice.MyRequest) (r *myservice.MyResponse, err error) {
	return &myservice.MyResponse{
		Res: req.Req,
	}, nil
}

func createClient(addr string) (c myservice.MyService, close func(), err error) {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	var transport thrift.TTransport

	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		return nil, nil, err
	}

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return nil, nil, err
	}

	err = transport.Open()
	if err != nil {
		return nil, nil, err
	}

	c = myservice.NewMyServiceClientFactory(transport, protocolFactory)

	close = func() {
		_ = transport.Close()
	}

	return c, close, nil
}

func createServer(addr string, handler myservice.MyService) (func(), error) {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return nil, err
	}

	processor := myservice.NewMyServiceProcessor(handler)

	server := thriftutils.NewServer4(processor, transport, transportFactory, protocolFactory)

	if err := server.Listen(); err != nil {
		return nil, err
	}

	go server.AcceptLoop()

	return func() {
		_ = server.Stop()
		_ = transport.Close()
	}, nil
}

type handler struct {
	ping        func(ctx context.Context) (err error)
	getResponse func(ctx context.Context, id myservice.UUID, req *myservice.MyRequest) (r *myservice.MyResponse, err error)
}

func (h *handler) Ping(ctx context.Context) (err error) {
	return h.ping(ctx)
}

func (h *handler) GetResponse(ctx context.Context, id myservice.UUID, req *myservice.MyRequest) (r *myservice.MyResponse, err error) {
	return h.getResponse(ctx, id, req)
}
