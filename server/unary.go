package main

import (
	"context"

	pb "github.com/Adilfarooque/grpc_demo_streaming/pb"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam)(*pb.HelloResponse,error){
	return &pb.HelloResponse{
		Message: "Hello",
	},nil
}