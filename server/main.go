package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/mertcaliskanlnx/golang-grpc-v1/server/pb"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, &srv{})

	fmt.Println("GRPC - Server V1.0.0 is   running...")
	err = server.Serve(lis)
	log.Fatal(err)
}

type srv struct{}

func (s *srv) Now(ctx context.Context, req *pb.NowRequest) (*pb.NowResponse, error) {
	return &pb.TimeResponse{Time: &pb.Time{
		Value: time.Now().String(),
	}}, nil
}

func (s *srv) Stream(req *pb.StreamRequest, stream pb.UserService_StreamServer) error {
	deadline := time.Now().Add(time.Duration(req.Length()) * time.Second)

	for !time.Now().After(deadline) {
		stream.Send(&pb.TimeResponse{Time: &pb.Time{
			Value: time.Now().String(),
		}})
		time.Sleep(1 * time.Second)
	}
	return nil
}
