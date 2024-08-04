package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"grpc-gateway/pkg/pb"
	"log"
	"net/http"

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
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &calculatorService{})
	pb.RegisterSampleServiceServer(s, &sampleService{})

	gwmux := runtime.NewServeMux()
	gwServer := &http.Server{
		Addr:    ":8091",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8091")
	log.Fatalln(gwServer.ListenAndServe())

}

func main() {
	StartGrpcServer()
}
