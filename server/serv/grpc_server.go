package serv

import (
	"fmt"
	"log"
	"net"

	"github.com/kolharsam/goki/gokigrpc"
	"google.golang.org/grpc"
)

func runServer() {
	fmt.Println("starting the goki server...")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen to: %v", err)
	}

	servInstance := gokigrpc.Server{}

	grpcServer := grpc.NewServer()

	gokigrpc.RegisterGokiGRPCServer(grpcServer, &servInstance)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
