package main

import (
	"context"
	"log"

	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)
	clientCreateBlog(client)
}

func clientCreateBlog(c pb.BlogServiceClient) string {
	log.Println("----- createBlog was invoked ------")

	blog := &pb.Blog{
		AuthorId: "Moitinho",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("unexpected error: %v\n", err)
	}

	log.Printf("blog was created: %v", res.Id)
	return res.Id
}
