package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"os"
	"vientiane/http/router"
)

func main() {
	// 初始化 zap 日志
	vlog := zap.NewExample()
	defer vlog.Sync()

	// 记录 gin 日志到文件。
	f, _ := os.Create("./http/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 同时将日志写到文件及控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	// TODO module api 各个模块的 API 写在这里
	router.HandleHealth(r, "health")
	router.HandleAccount(r, "account")

	//flag.Parse()
	r.Run(":8080")
}
