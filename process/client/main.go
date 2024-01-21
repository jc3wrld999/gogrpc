package main

import (
	"context"
	"fmt"
	pb "github.com/jc3wrld999/gogrpc/process/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	address := "localhost:50001"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProcessServiceClient(conn)

	command := "ls -l"
	response, err := client.ExecuteCommand(context.Background(), &pb.CommandRequest{Command: command})
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	fmt.Printf("Command Output:\n%s\n", response.Output)
}
