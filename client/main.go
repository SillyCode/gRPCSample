package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "grpc/sample1/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGet_CarClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCar(ctx, &pb.CAR{
		Brand:           "Tesla",
		Model:           "Model3",
		NumberSeats:     4,
		EngineType:      pb.ENGINE_TYPE_EV,
		NumberOfAirbags: 4,
		Style:           pb.Style_HIGH,
	})

	if err != nil {
		log.Fatalf("could not get new car: %v", err)
	}
	log.Printf("New Car: %s", r.GetMessage())
}
