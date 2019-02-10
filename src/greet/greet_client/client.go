package main

import (
	"fmt"
	"log"

	"github.com/gRPC-go-microservices/src/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	// by default gRPC has SSL, for now, without this
	conn, err := grpc.Dial("localhost:500051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// defer means defer this statement at the very end of this function
	defer conn.Close()

	// create a new client
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client: %f", c)
}
