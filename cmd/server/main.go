package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"turbo/pkg/adder"
	"turbo/pkg/api"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:4444")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	s := grpc.NewServer()
	srv := &adder.GRPCServer{UnimplementedAdderServer: api.UnimplementedAdderServer{}}
	api.RegisterAdderServer(s, srv)

	log.Fatalln(s.Serve(listen))
}
