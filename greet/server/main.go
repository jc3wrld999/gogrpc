package main

import (
	pb "github.com/jc3wrld999/gogrpc/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50001"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Listening on %s\n", addr)
	}
	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v\n", err)
	}
}
