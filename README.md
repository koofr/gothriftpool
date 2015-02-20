gothriftpool
============

Go Thrift pool proxy generator.

# Example

```
go get github.com/koofr/gothriftpool/gothriftpool
thrift --gen go:thrift_import=github.com/koofr/thrift/lib/go/thrift -out $GOPATH/src gothriftpooltest/myservice.thrift
gothriftpool -w myservice.MyService
```

The latter will create `$GOPATH/src/myserviceproxy/myserviceproxy.go` file.

# Testing

```
go test github.com/koofr/gothriftpool/gothriftpooltest
```
