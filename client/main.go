package main

import (
	"context"
	"io"
	"log"

	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
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
	//clientReadBlog(client)
	//clientUpdateBlog(client, "631a9c0bcfaef11492da3955")
	listBlogClient(client)
	deleteBlogClient(client, "631a9c0bcfaef11492da3955")
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
		AuthorId: "Teste 3",
		Title:    "My Third Blog",
		Content:  "Content of the Third blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("unexpected error: %v\n", err)
	}

	log.Printf("blog was created: %v", res.Id)
	return res.Id
}

func clientUpdateBlog(c pb.BlogServiceClient, id string) {
	log.Println("----- clientUpdateBlog was invoked ------")

	blog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Luiz",
		Title:    "A new title",
		Content:  "Content of the fisrt blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("error happened while updating: %v \n", err)
	}
}

func listBlogClient(c pb.BlogServiceClient) {
	log.Println("----- listBlogClient was invoked ------")
	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error while calling ListBlog: %v \n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("something happened: %v", err)
		}

		log.Println(res)
	}
}
func deleteBlogClient(c pb.BlogServiceClient, id string) {
	log.Println("----- deleteBlogClient was invoked ------")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("error while deleting: %v\n", err)
	}
	log.Println("blog was deleted")

}
