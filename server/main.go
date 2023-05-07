package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/mertcaliskanlnx/golang-grpc-v1/server/pb"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	// Proto'da tanimladigimiz TimeService servisini implement etmek istedigimizi bu sekilde belirtiyoruz.
	pb.RegisterTimeServiceServer(server, &srv{})
	log.Printf("starting gRPC server")
	err = server.Serve(lis)
	log.Fatal(err)
}

type srv struct{}

func (s *srv) Now(ctx context.Context, req *pb.NowRequest) (*pb.TimeUpdate, error) {
	return &pb.TimeUpdate{Time: &pb.Time{
		Value: time.Now().String(),
	}}, nil
}

func (s *srv) Stream(req *pb.TimeStreamRequest, stream pb.TimeService_StreamServer) error {
	// bir deadline belirleyip sonuna kadar client'a veri gonderiyoruz (server stream)
	deadline := time.Now().Add(time.Duration(req.Length) * time.Second)
	for !time.Now().After(deadline) {
		time.Sleep(time.Millisecond * 300)

		stream.Send(&pb.TimeUpdate{Time: &pb.Time{
			Value: time.Now().String(),
		}})
	}
	return nil
}
