package main

import (
	"gin/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//设置运行模式
	//Gin支持三种模式
	//1.gin.DebugMode（调试模式）
	//2.gin.ReleaseMode（发布模式）
	//3.gin.TestMode（测试模式）
	gin.SetMode(gin.DebugMode)

	//路由初始化
	r := routes.InitRouter()
	err := r.Run()
	log.Fatalf("运行出错，错误为：%v\n", err)
}
