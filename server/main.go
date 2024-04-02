package main

import (
	"log"
	"net"

	pb "github.com/Adilfarooque/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer,&helloServer{})
	log.Printf("server started at %v",list.Addr())
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}
