package server

import (
	"context"
	"log"

	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"
	"github.com/luizmoitinho/api-grpc-mongodb/service"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog was invoked with %v\n", in)
	data, err := service.CreateBlog(ctx, in)
	return data, err
}
