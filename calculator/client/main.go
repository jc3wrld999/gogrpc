package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/jc3wrld999/gogrpc/calculator/proto"
)

var addr string = "localhost:50001"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	con := pb.NewCalculatorServiceClient(conn)
	doSum(con)
}
