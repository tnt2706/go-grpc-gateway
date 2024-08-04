package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-gateway/pkg/calculator_service"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	result := in.A() + in.GetSecondNumber()
	fmt.Println(result)
	return &pb.SumResponse{Result: result}, nil
}

func StartGrpcServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:6903")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})
	log.Printf("🚀 Running listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	StartGrpcServer()
}
