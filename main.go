package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/bks71/traintix/api_server"
	"github.com/bks71/traintix/pb"
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
	s := grpc.NewServer()
	pb.RegisterTicketingServer(s, &api_server.TicketingServer{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
