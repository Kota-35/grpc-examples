package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "example.com/grpc-examples/echo_simple/echo_simple"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The Server Port")
)

type echoServer struct {
	pb.UnimplementedEchoSimpleServiceServer
}

func (s *echoServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %q", req.GetMessage())
	return &pb.EchoResponse{Message: req.GetMessage()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoSimpleServiceServer(s, &echoServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
