package main

import (
	"log"
	"net"

	pb "github.com/Adilfarooque/grpc_demo_streaming/pb"
	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type helloServer struct{
	pb.GreetServiceServer
}

func main() {
	list, err := net.Listen("tcp",port)
	if err != nil{
		log.Fatalf("Failed to start the server %v",err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer,&helloServer{})
	if err := grpcServer.Serve(list); err != nil{
		log.Fatalf("Failed to start %v",err)
	}
}

