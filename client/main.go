package main

import (
	"log"

	pb "github.com/Adilfarooque/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	connection, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewGreetServiceClient(connection)

	names := &pb.NamesList{
		Names: []string{"Amjad", "Adil", "Ashfak", "Dilsha"},
	}

	//callSayHello(client)
	callSayHelloServerStream(client, names)
}
