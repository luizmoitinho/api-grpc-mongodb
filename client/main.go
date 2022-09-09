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
	//clientCreateBlog(client)
	clientReadBlog(client)
}

func clientReadBlog(c pb.BlogServiceClient) *pb.Blog {
	log.Println("----- clientReadBlog was invoked ------")

	req := &pb.BlogId{
		Id: "631a9c0bcfaef11492da3a955",
	}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("unexpected error: %v\n", err)
	}

	log.Printf("blog was read: %v", res)
	return res
}

func clientCreateBlog(c pb.BlogServiceClient) string {
	log.Println("----- clientCreateBlog was invoked ------")

	blog := &pb.Blog{
		AuthorId: "Moitinho2",
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
