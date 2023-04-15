package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"test-grpc/pb"
	"time"
)

const (
	address     = "localhost:9000"
	defaultName = "dr-who"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewPingPongServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.PingPong(ctx, &pb.PingRequest{PingEntry: &pb.Ping{Data: "PING"}})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Println(r.GetResult())
}
