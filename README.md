# traintix

Makefile commands:

```shell
protoc --go_out=. --go-grpc_out=. pb/ticketing.proto
go run server/main.go [--port 5280]
staticcheck ./...
mockgen -source=internal/inventory/api.go -package="inventory" -destination=internal/inventory/inventory_mock.go
gofmt -d -s client/ server/
go test ./server/...
goimports -w client/ server/
```

TODO

* makefile
* test client
* finish API methods
* Postman test scripts
