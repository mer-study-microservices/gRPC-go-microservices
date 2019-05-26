package main

import (
	"fmt"
	"log"

	"github.com/gRPC-go-microservices/setup/client-boilerplate/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello I'm a client")
	conn, err := grpc.Dial("localhost:50551", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	// this gets called at the very end
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client: %f", c)
}
