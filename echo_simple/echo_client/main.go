package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example.com/grpc-examples/echo_simple/echo_simple"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoSimpleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.EchoRequest{Message: "こんにちは! I'm Kota Ishikuro"})
	if err != nil {
		log.Fatalf("Could not echo: %v", err)
	}
	log.Printf("Echo: %s", r.GetMessage())
}
