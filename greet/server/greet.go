package main

import (
	"context"
	pb "github.com/jc3wrld999/gogrpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet Function was invoked with %v\n", request)
	return &pb.GreetResponse{
		Result: "Hello " + request.FirstName,
	}, nil
}
