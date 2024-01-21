package main

import (
	"context"
	pb "github.com/jc3wrld999/gogrpc/calculator/proto"
	"log"
)

func doSum(client pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstName:  1,
		SecondName: 1,
	})

	if err != nil {
		log.Fatalf("Could not sum %v\n", err)
	}
	log.Printf("Sum %s\n", res.Result)
}
