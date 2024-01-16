package server

import (
	"log"

	"github.com/bks71/traintix/internal/api"
	"google.golang.org/grpc"
)

func NewGRPCServer() *grpc.Server {
	//create grpc server
	grpcServer := grpc.NewServer()

	//register Ticketing API server
	if ts, err := api.NewTicketingServer(); err == nil {
		ts.RegisterService(grpcServer)
	} else {
		log.Fatalf("unable to register TicketingServer %v", err)
	}

	return grpcServer
}
