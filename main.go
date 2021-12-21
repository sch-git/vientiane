package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"vientiane/pub/adapter"
	vientiane "vientiane/pub/idl/grpc"
	"vientiane/server/controller"
)

func main() {
	grpcServer := grpc.NewServer()
	vientiane.RegisterVientianeServiceServer(grpcServer, new(controller.VientianeServiceImpl))

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", adapter.Port))
	if nil != err {
		log.Fatalln(err)
	}

	err = grpcServer.Serve(listen)
	if nil != err {
		log.Fatalln(err)
	}
}
