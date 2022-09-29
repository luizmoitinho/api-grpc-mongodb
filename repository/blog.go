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

func List(ctx context.Context, stream pb.BlogService_ListBlogServer) error {
	cursor, err := collection.Find(ctx, primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknow internal error: %v", err),
		)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		data := &types.BlogItem{}
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("error while deconding data from mongoDB: %v", err),
			)
		}
		stream.Send(types.DocumentToBlog(data))
	}

	if err = cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("error while deconding data from mongoDB: %v", err),
		)
	}

	return nil
}

func Delete(ctx context.Context, oid primitive.ObjectID) error {
	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot delete object in mongoDB with id: %v", oid),
		)
	}

	if res.DeletedCount == 0 {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("blog was not found, id: %v", oid),
		)
	}

	return nil
}
