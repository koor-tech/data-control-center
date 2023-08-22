package server

import (
	"github.com/koor-tech/data-control-center/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server.RegisterGRPCHandler(grpcServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
