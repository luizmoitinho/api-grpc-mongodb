package server

import (
	"context"
	"log"

	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"
	"github.com/luizmoitinho/api-grpc-mongodb/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog was invoked with %v\n", in)
	data, err := service.CreateBlog(ctx, in)
	return data, err
}

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with %v\n", in)
	data, err := service.ReadBlog(ctx, in)
	return data, err
}

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked with %v\n", in)
	data, err := service.UpdateBlog(ctx, in)
	return data, err
}

func (s *Server) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Printf("ListBlog was invoked with %v\n", in)
	err := service.ListBlog(in, stream)
	return err
}

func (s *Server) DeleteBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with %v\n", in)
	return service.DeleteBlog(ctx, in)
}
