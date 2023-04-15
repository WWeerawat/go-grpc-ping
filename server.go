package main

import (
	"context"
	"fmt"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"log"
	"net"
	"test-grpc/pb"
)

type PingPongServer struct {
	pb.UnimplementedPingPongServiceServer
}

func (p *PingPongServer) PingPong(ctx context.Context, req *pb.PingRequest) (*pb.PongResponse, error) {
	res := &pb.PongResponse{}
	if req == nil {
		fmt.Println("request must not be nil")
		return res, xerrors.Errorf("request must not be nil")
	}
	res = &pb.PongResponse{Result: "PONG  "}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "9000"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterPingPongServiceServer(srv, &PingPongServer{})
	log.Printf("Server start at port %s", "9000")

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
