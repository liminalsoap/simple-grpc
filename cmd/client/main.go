package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"turbo/pkg/api"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:4444", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	c := api.NewAdderClient(conn)
	var a, b int32 = 1, 2
	res, err := c.Add(context.Background(), &api.AddRequest{X: a, Y: b})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}
