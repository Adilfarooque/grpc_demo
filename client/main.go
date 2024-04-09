package main

import (
	"log"

	pb "github.com/Adilfarooque/grpc_demo_streaming/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NameList{
	// 	Names: []string{"Amjad", "Adil", "Ashfak", "Dilsha"},
	// }

	callsayHello(client)
}
