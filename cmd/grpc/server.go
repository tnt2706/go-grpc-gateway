package grpc

import (
	"context"
	"fmt"
	"grpc-gateway/pkg/generate/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type calculatorService struct {
	pb.UnimplementedCalculatorServiceServer
}

type sampleService struct {
	pb.UnimplementedSampleServiceServer
}

func (s *calculatorService) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	result := in.Number_1 + in.Number_2
	fmt.Println(result)
	return &pb.SumResponse{Result: result}, nil
}

func (s *sampleService) Print(ctx context.Context, in *pb.PrintRequest) (*pb.PrintResponse, error) {
	return &pb.PrintResponse{Result: "Hello"}, nil
}

func StartGrpcServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:6903")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &calculatorService{})
	pb.RegisterSampleServiceServer(s, &sampleService{})
	log.Printf("🚀 Running listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	StartGrpcServer()
}
