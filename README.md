# traintix

Makefile commands:

```shell
protoc --go_out=. --go-grpc_out=. pb/ticketing.proto
go run server/main.go [--port 5280]
staticcheck ./...
mockgen -source=server/inv/api.go -package="inv" -destination=server/inv/inventory_mock.go
gofmt -d -s client/ server/
go test ./server/...
goimports -w client/ server/
```

TODO

* makefile
* implement auth
* test client
* finish API methods
* Postman test scripts
