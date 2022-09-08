package main

import (
	"log"
	"net"

	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"

	"github.com/luizmoitinho/api-grpc-mongodb/repository"
	"github.com/luizmoitinho/api-grpc-mongodb/server"
	"google.golang.org/grpc"
)

func main() {
	var addr string = "0.0.0.0:50051"

	repository.MongoConfigCollection()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, &server.Server{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}

}
