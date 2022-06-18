package main

import (
	"log"
	"net"

	"github.com/cory-evans/Ichnaea/pkg/position"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	// create a new grpc server
	s := grpc.NewServer()
	position.RegisterService(s)

	// start the server
	log.Println("Starting the server on port 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
