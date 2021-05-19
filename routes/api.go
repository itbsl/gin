package routes

import (
	"gin/app/controllers/demo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	//指定模板路径
	r.LoadHTMLGlob("app/views/**/*")
	//指定静态文件路径
	//第一个参数的名字是在静态文件里引用时的起始路径名，第二个参数是真实文件名
	r.Static("/s", "./static")

	//注册路由
	v1 := r.Group("/v1")
	{
		demoGroup := v1.Group("/demo")
		{//大括号没实际意义，只是为了让每一组看起来更像是一组
			demoGroup.GET("/index", demo.Index)
			demoGroup.GET("/ping", demo.Ping)
			demoGroup.GET("/GetParam/:name", demo.GetParam)
			demoGroup.GET("/GetQuery", demo.GetQuery)
			demoGroup.POST("/PostExample", demo.Post)
		}
	}

	//Not Found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg": "接口不存在",
		})
	})
	return r
}
