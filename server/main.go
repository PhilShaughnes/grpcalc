package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/PhilShaughnes/grpcalc/pb"
)

const port = ":8080"

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	log.Printf("Got Add request: {a:%v, b:%v}", in.A, in.B)
	return &pb.CalculationResponse{
		Result: in.A + in.B,
	}, nil
}

func (s *server) Divide(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	log.Printf("Got Divide request: {a:%v, b:%v}", in.A, in.B)
	if in.B == 0 {
		return nil, status.Error(codes.InvalidArgument, "Cannot divide by zero")
	}

	return &pb.CalculationResponse{
		Result: in.A / in.B,
	}, nil
}

func (s *server) Sum(
	ctx context.Context,
	in *pb.NumbersRequest,
) (*pb.CalculationResponse, error) {
	var sum int64
	log.Printf("Got Sum request: {Numbers:%v}", in.Numbers)

	for _, num := range in.Numbers {
		sum += num
	}
	return &pb.CalculationResponse{
		Result: sum,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterCalculatorServer(s, &server{})
	log.Printf("listening on: %v", port)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
