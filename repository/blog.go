package repository

import (
	"context"
	"fmt"

	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/luizmoitinho/api-grpc-mongodb/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOne(ctx context.Context, data types.BlogItem) (*pb.BlogId, error) {
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot convert to OID: %v", err),
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
