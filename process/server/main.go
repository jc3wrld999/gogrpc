package main

import (
	"context"
	"log"
	"os/exec"

	pb "github.com/jc3wrld999/gogrpc/process/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	pb.ProcessServiceServer
}

func (s *Server) ExecuteCommand(ctx context.Context, req *pb.CommandRequest) (*pb.CommandResponse, error) {
	cmd := exec.Command("sh", "-c", req.Command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return &pb.CommandResponse{
		Output: string(output),
	}, nil
}

func main() {
	var addr string = "0.0.0.0:50001"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", addr, err)
	}

	s := grpc.NewServer()
	pb.RegisterProcessServiceServer(s, &Server{})

	log.Printf("Server listening on %s", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
