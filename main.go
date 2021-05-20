package main

import (
	"errors"
	"gin/extend/setting"
	"gin/global"
	"gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func init() {
	//设置运行模式
	//Gin支持三种模式
	//1.gin.DebugMode（调试模式）
	//2.gin.ReleaseMode（发布模式）
	//3.gin.TestMode（测试模式）
	gin.SetMode(gin.TestMode)

	//加载配置文件
	err := loadConfig()
	if err != nil {
		log.Fatalf("读取配置文件出错，错误为：%v\n", err)
	}
}

func main() {
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


// loadConfig 加载配置文件
func loadConfig() error {
	s, err := setting.NewSetting()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			return errors.New("配置文件未找到")
		} else {
			// 配置文件被找到，但产生了另外的错误
			return err
		}
	}
	err = s.ReadSection("server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = s.ReadSection("app", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("log", &global.LogSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("mysql", &global.MySQLSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("rabbitmq", &global.RabbiMQSetting)
	if err != nil {
		return err
	}

	return nil
}
