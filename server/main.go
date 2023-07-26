package main

import (
	"context"
	"log"
	"net"

	pb "grpc/sample1/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGet_CarServer
}

func (s *server) GetCar(ctx context.Context, in *pb.CAR) (*pb.GetCarResponse, error) {
	log.Printf("Received: %v", in.Brand)
	return &pb.GetCarResponse{Message: "Recieved car brand " + in.Brand}, nil
}

func main() {
	port := ":50051"
	srv := grpc.NewServer()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterGet_CarServer(srv, &server{})

	log.Printf("Listening on %s", port)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
