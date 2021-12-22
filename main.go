package main

import (
	"fmt"
<<<<<<< HEAD
	"github.com/golang/glog"
	"google.golang.org/grpc"
=======
	"google.golang.org/grpc"
	"log"
>>>>>>> master
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
<<<<<<< HEAD
		glog.Fatalln(err)
=======
		log.Fatalln(err)
>>>>>>> master
	}

	err = grpcServer.Serve(listen)
	if nil != err {
<<<<<<< HEAD
		glog.Fatalln(err)
=======
		log.Fatalln(err)
>>>>>>> master
	}
}
