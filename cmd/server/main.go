package main

import (
	"context"
	"log"
	"net"

	"github.com/franklaercio/grpc-greeting/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server é usado para implementar example.MyService.
type server struct {
	pb.UnimplementedGreetingServer
}

// GreetingResponse implementa SayHello
func (s *server) SayHello(ctx context.Context, in *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	log.Printf("Received message body from client: %s", in.Data)
	return &pb.GreetingResponse{Result: "Greeting from the server!"}, nil
}

func main() {
	println("Starting new gRPC server...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterGreetingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
