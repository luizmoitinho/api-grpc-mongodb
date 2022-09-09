package service

import (
	"context"

	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"
	"github.com/luizmoitinho/api-grpc-mongodb/repository"
	"github.com/luizmoitinho/api-grpc-mongodb/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"cannot parse ID",
		)
	}
	res, err := repository.Get(ctx, oid)
	return types.DocumentToBlog(res), err
}

func UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"cannot parse ID",
		)
	}

	data := &types.BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	err = repository.Update(ctx, oid, data)
	return &emptypb.Empty{}, err
}
