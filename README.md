# traintix

```shell
protoc --go_out=. --go-grpc_out=. pb/ticketing.proto
go run cmd/traintix.go
go run cmd/traintix.go --port 6789
```
