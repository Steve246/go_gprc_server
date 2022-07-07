package main

import (
	"fmt"
	pb "go_gprc_server/api"
	"go_gprc_server/calc"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	host := os.Getenv("GPRC_HOST") // set GPRC_HOST=localhost
	port := os.Getenv("GPRC_PORT") // set GPRC_PORT=8888
	serverInfo := fmt.Sprintf("%s:%s", host, port)

	listen, err := net.Listen("tcp", serverInfo)

	if err != nil {
		log.Fatalln("failed to listed", err)
	}

	s := calc.Server{}
	grpcServer := grpc.NewServer()
	
	pb.RegisterCalculatorServer(grpcServer, &s)

	log.Fatalln("Server run on", serverInfo)

	err = grpcServer.Serve(listen)

	if err != nil {
		log.Fatalln("failed to server", err)
	}
}