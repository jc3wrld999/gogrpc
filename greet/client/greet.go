package main

import (
	"context"
	pb "github.com/jc3wrld999/gogrpc/greet/proto"
	"log"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "jc3wrld999",
	})

	if err != nil {
		log.Fatalf("Could not Greet %v\n", err)
	}
	log.Printf("Greeting %s\n", res.Result)
}
