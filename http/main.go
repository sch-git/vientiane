package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/web"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
	"vientiane/http/router"
	"vientiane/http/utils"
)

func main() {
	// 初始化 zap 日志
	vlog := zap.NewExample()
	defer vlog.Sync()
	defer utils.KafkaClose()

	// 记录 gin 日志到文件。
	f, _ := os.Create("./http/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 同时将日志写到文件及控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	// TODO module api 各个模块的 API 写在这里
	router.HandleHealth(r, "health")
	router.HandleAccount(r, "account")
	router.HandleArticle(r, "article")
	router.HandleIndices(r, "indices")

	//flag.Parse()
	//r.Run(":8080")
	go func() {
		m := gin.Default()
		router.HandleMonitor(m, "metrics")
		monitorService := web.NewService(web.Address(":9910"), web.Handler(m))
		err := monitorService.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		redisCli, clean := utils.NewRedisClient(context.Background())
		defer clean()
		redisCli.Set(context.Background(), "vientiane", "hello world!", 0)
	}()

	microService := web.NewService(
		web.Address(":8008"),
		web.Handler(r),
	)
	microService.Run()
}
