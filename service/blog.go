package service

import (
	"context"

	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"
	"github.com/luizmoitinho/api-grpc-mongodb/repository"
	"github.com/luizmoitinho/api-grpc-mongodb/types"
)

func CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	data := types.BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := repository.InsertOne(ctx, data)
	return res, err
}
