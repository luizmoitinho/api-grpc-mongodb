package repository

import (
	"context"
	"fmt"

	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/luizmoitinho/api-grpc-mongodb/types"
	"go.mongodb.org/mongo-driver/bson"
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

func Get(ctx context.Context, oid primitive.ObjectID) (*types.BlogItem, error) {
	data := &types.BlogItem{}

	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"cannot find blog with the id provided",
		)
	}

	return data, nil
}

func Update(ctx context.Context, oid primitive.ObjectID, data *types.BlogItem) error {

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return status.Errorf(
			codes.Internal,
			"cannot not update",
		)
	}

	if res.MatchedCount == 0 {
		return status.Errorf(
			codes.NotFound,
			"cannot find blog with id",
		)
	}

	return nil
}
