package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Adilfarooque/grpc_demo_streaming/pb"
)

func callsayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not great: %v", err)
	}
	log.Printf("%s", res.Message)

}
