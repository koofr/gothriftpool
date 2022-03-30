gothriftpool
============

Go Thrift pool proxy generator.

# Example

```sh
go install github.com/koofr/gothriftpool/gothriftpool
thrift --gen go -out . gothriftpooltest/myservice.thrift
gothriftpool -w myservice.MyService
```

The latter will create `./myserviceproxy/myserviceproxy.go` file.

# Testing

```sh
go test ./gothriftpooltest
```

Coverage:

```sh
cd gothriftpooltest
go test -coverpkg=./myserviceproxy -coverprofile=test.coverprofile && go tool cover -html=test.coverprofile
```

## Regenerate test service

```sh
thrift --gen go -out `pwd`/gothriftpooltest gothriftpooltest/myservice.thrift
go run ./gothriftpool myservice.MyService > gothriftpooltest/myserviceproxy/myserviceproxy.go
```
