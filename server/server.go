package server

import (
	pb "github.com/luizmoitinho/api-grpc-mongodb/proto"
)

type Server struct {
	pb.BlogServiceServer
}
