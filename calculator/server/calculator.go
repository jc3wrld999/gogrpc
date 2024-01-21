package main

import (
	"context"
	pb "github.com/jc3wrld999/gogrpc/calculator/proto"
	"log"
)

func (s *Server) Calculator(ctx context.Context, request *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Calculator Function was invoked with %v\n", request)
	return &pb.SumResponse{
		Result: request.FirstName + request.SecondName,
	}, nil
}
