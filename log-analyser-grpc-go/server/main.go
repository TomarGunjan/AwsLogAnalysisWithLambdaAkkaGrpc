package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log-analyser-grpc-go/helper"
	"log-analyser-grpc-go/proto"
	"log-analyser-grpc-go/services"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// This function sets up the server at port 50051
func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := services.NewLogAnalyserService(helper.GetHelper())
	proto.RegisterLogAnalyserServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
