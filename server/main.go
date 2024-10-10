package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("This will be a sample gRPC server entrypoint.")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	log.Printf("gRPC server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", port, err)
	}
}
