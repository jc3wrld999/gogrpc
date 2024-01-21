package main

import (
	"fmt"
	pb "github.com/jc3wrld999/gogrpc/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(request *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {

	log.Printf("Greet Many Times function was invoked with: %v\n", request)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", request.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
