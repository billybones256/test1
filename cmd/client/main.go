package main

import (
	"context"
	"flag"
	"log"

	"task1/pkg/api"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	connection, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewCheckerClient(connection)
	res, err := client.Check(context.Background(), &api.CheckRequest{Message: flag.Arg(0)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetAnswer())
}
