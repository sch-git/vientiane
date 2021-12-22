package main

import (
	"fmt"
	"github.com/golang/glog"
	"google.golang.org/grpc"
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
		glog.Fatalln(err)
	}

	err = grpcServer.Serve(listen)
	if nil != err {
		glog.Fatalln(err)
	}
}
