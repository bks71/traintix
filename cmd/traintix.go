package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/bks71/traintix/internal/server"
	"google.golang.org/grpc/reflection"
)

var port = flag.Int("port", 5280, "server port")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := server.NewGRPCServer()

	// Register reflection service on gRPC server.
	reflection.Register(grpcSrv)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
