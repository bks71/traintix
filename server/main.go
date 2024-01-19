package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/bks71/traintix/server/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port = flag.Int("port", 5280, "server port")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create grpc server
	grpcServer := grpc.NewServer()

	//register Ticketing API server
	if ts, err := api.NewTicketingServer(); err == nil {
		ts.RegisterService(grpcServer)
	} else {
		log.Fatalf("unable to register TicketingServer %v", err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
