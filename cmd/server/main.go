package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"task1/pkg/api"
	"task1/pkg/checker"
)

func main() {
	server := grpc.NewServer()
	microService := &checker.GRPCServer{}
	api.RegisterCheckerServer(server, microService)
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
