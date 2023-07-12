package main

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/web"
	"go.uber.org/zap"
	"io"
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
	router.HandleArticle(r, "indices")

	//flag.Parse()
	//r.Run(":8080")

	microService := web.NewService(
		web.Address(":8088"),
		web.Handler(r),
	)
	microService.Run()
}
