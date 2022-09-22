package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"vientiane/pub/adapter"
	vientiane "vientiane/pub/idl/grpc"
	"vientiane/server/controller"
	"vientiane/server/mq"
)

func main() {
	// 初始化 zap 日志
	vlog := zap.NewExample()
	defer vlog.Sync()
	vlog.Info("server start ...")

	// 初始化 grpc 服务端
	grpcServer := grpc.NewServer()
	vientiane.RegisterVientianeServiceServer(grpcServer, new(controller.VientianeServiceImpl))

	// 启动 article 消息监听
	go mq.ArticleConsumer()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", adapter.Port))
	if nil != err {
		vlog.Fatal("net listen err", zap.Error(err))
	}

	err = grpcServer.Serve(listen)
	if nil != err {
		vlog.Fatal("serve err", zap.Error(err))
	}
}
