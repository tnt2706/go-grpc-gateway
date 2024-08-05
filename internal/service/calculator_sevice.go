package service

import (
	"context"
	"grpc-gateway/pkg/pb"
)

type CalculatorService struct {
	pb.UnimplementedGreeterServer
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (us *CalculatorService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: in.Name + " world"}, nil
}
