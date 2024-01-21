Server code lives here.

To start the server:

```shell
go run server/main.go [--port 5280]
```


sequenceDiagram
   GRPC_client->>api/TicketingServer: Purchase
    api/TicketingServer->>inventory/Inventory: ReserveSeat
    inventory/Inventory-->>api/TicketingServer: Reservation
    api/TicketingServer-->>GRPC_client: Receipt
