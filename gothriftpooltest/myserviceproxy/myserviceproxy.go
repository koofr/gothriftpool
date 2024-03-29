// Autogenerated by gothriftpool generator
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package myserviceproxy

import (
	"context"

	"github.com/koofr/go-resourcepool"

	"github.com/koofr/gothriftpool/gothriftpooltest/myservice"
)

type ErrCreateResource struct {
	Err error
}

func NewErrCreateResource(err error) *ErrCreateResource {
	return &ErrCreateResource{Err: err}
}

func (e *ErrCreateResource) Error() string {
	return "myservice.MyService resource could not be created: " + e.Err.Error()
}

func (e *ErrCreateResource) Unwrap() error {
	return e.Err
}

type ClientFactory func(context.Context) (client myservice.MyService, close func(), err error)

type resource struct {
	client   myservice.MyService
	close    func()
	isClosed bool
}

func (r *resource) Close() {
	r.isClosed = true
	r.close()
}

func (r *resource) IsClosed() bool {
	return r.isClosed
}

type ClientPool struct {
	clientFactory ClientFactory
	pool          *resourcepool.ResourcePool
	retries       int
}

func NewClientPool(clientFactory ClientFactory, idleCapacity int, maxResources int) *ClientPool {
	p := &ClientPool{
		clientFactory: clientFactory,
		retries:       2,
	}

	p.pool = resourcepool.NewResourcePool(p.createResource, idleCapacity, maxResources)

	return p
}

func (p *ClientPool) Close() {
	p.pool.Close()
}

func (p *ClientPool) SetRetries(retries int) {
	p.retries = retries
}

func (p *ClientPool) GetPool() *resourcepool.ResourcePool {
	return p.pool
}

func (p *ClientPool) createResource(ctx context.Context) (r resourcepool.Resource, err error) {
	client, close, err := p.clientFactory(ctx)
	if err != nil {
		return nil, err
	}

	r = &resource{
		client:   client,
		close:    close,
		isClosed: false,
	}

	return
}

func (p *ClientPool) GetClient(ctx context.Context) (client myservice.MyService, release func(), err error) {
	var lastErr error

	for i := 0; i < p.retries; i++ {
		r, err := p.pool.Acquire(ctx)
		if err != nil {
			if err == context.Canceled || err == context.DeadlineExceeded {
				return nil, nil, err
			}

			p.pool.Empty()

			lastErr = err
			continue
		}

		err = r.(*resource).client.Ping(ctx)
		if err != nil {
			r.Close()

			p.pool.Release(r)

			p.pool.Empty()

			lastErr = err
			continue
		}

		release = func() {
			if ctx.Err() != nil {
				r.Close()
			}

			p.pool.Release(r)
		}

		return r.(*resource).client, release, nil
	}

	return nil, nil, NewErrCreateResource(lastErr)
}

type Proxy struct {
	clientFactory ClientFactory
}

func New(clientFactory ClientFactory) *Proxy {
	return &Proxy{
		clientFactory: clientFactory,
	}
}

func (p *Proxy) Ping(ctx context.Context) (err error) {
	thriftClient, thriftClientClose, err := p.clientFactory(ctx)
	if err != nil {
		return
	}
	defer thriftClientClose()

	return thriftClient.Ping(ctx)
}

func (p *Proxy) GetResponse(ctx context.Context, id myservice.UUID, req *myservice.MyRequest) (r *myservice.MyResponse, err error) {
	thriftClient, thriftClientClose, err := p.clientFactory(ctx)
	if err != nil {
		return
	}
	defer thriftClientClose()

	return thriftClient.GetResponse(ctx, id, req)
}
