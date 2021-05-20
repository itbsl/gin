package main

import (
	"gin/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	//设置运行模式
	//Gin支持三种模式
	//1.gin.DebugMode（调试模式）
	//2.gin.ReleaseMode（发布模式）
	//3.gin.TestMode（测试模式）
	gin.SetMode(gin.DebugMode)

	//路由初始化
	router := routes.InitRouter()
	server := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	err := server.ListenAndServe()
	log.Fatalf("运行出错，错误为：%v\n", err)
}
