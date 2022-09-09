package server

import (
	"context"
	"log"

	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"
	"github.com/luizmoitinho/api-grpc-mongodb/service"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with %v\n", in)
	data, err := service.ReadBlog(ctx, in)
	return data, err
}
