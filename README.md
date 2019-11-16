gothriftpool
============

Go Thrift pool proxy generator.

# Example

```sh
go get github.com/koofr/gothriftpool/gothriftpool
thrift --gen go -out $GOPATH/src gothriftpooltest/myservice.thrift
gothriftpool -w myservice.MyService
```

The latter will create `$GOPATH/src/myserviceproxy/myserviceproxy.go` file.

# Testing

```sh
go test github.com/koofr/gothriftpool/gothriftpooltest
```

## Regenerate test service

```sh
thrift --gen go -out `pwd`/gothriftpooltest gothriftpooltest/myservice.thrift
go run gothriftpool/main.go myservice.MyService > gothriftpooltest/myserviceproxy/myserviceproxy.go
```
