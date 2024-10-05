package main

import (
	"fmt"
	"log"
	"net"

	cpb "github.com/SoumyadipPayra/protobufs/go-protos/cpp_compiler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "10001"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	fmt.Println("listening on port", port)

	// Create a new gRPC server instance.
	grpcServer := grpc.NewServer()

	// Register the ExampleService server implementation.
	cpb.RegisterCPPCompilerServer(grpcServer, &CPPCompilerServer{})

	// Register reflection service on gRPC server (useful for debugging and grpcurl).
	reflection.Register(grpcServer)

	defer grpcServer.GracefulStop()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
