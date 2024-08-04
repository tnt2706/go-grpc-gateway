package grpc

import (
	"context"
	"log"
	"time"

	pb "assignment/pkg/calculator/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func sum(client pb.CalculatorServiceClient, res *pb.SumRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := client.Sum(ctx, res)
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}

	log.Println(result)
}

func NewClient(addr string, opts ...grpc.DialOption) pb.CalculatorServiceClient {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := pb.NewCalculatorServiceClient(conn)
	return client
}

func SumTwo() {

	conn := NewClient("0.0.0.0:6903", grpc.WithTransportCredentials(insecure.NewCredentials()))

	sum(conn, &pb.SumRequest{FirstNumber: 10, SecondNumber: 20})
	sum(conn, &pb.SumRequest{FirstNumber: 20, SecondNumber: 30})
}
