package routes

import (
	"gin/app/controllers/article"
	"gin/app/controllers/demo"
	"gin/app/controllers/tag"
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
		//文章
		v1.GET("/articles", article.List)
		v1.GET("/articles/:id", article.Get)
		v1.POST("/articles", article.Create)
		v1.PUT("/articles/:id", article.Update)
		v1.DELETE("/articles/:id", article.Delete)
		v1.PATCH("/articles/:id/state", article.Update)

		//标签
		tagController := tag.NewTag()
		v1.GET("/tags", tagController.List)
		v1.GET("/tags/:id", tagController.Get)
		v1.POST("/tags", tagController.Create)
		v1.PUT("/tags/:id", tagController.Update)
		v1.DELETE("/tags/:id", tagController.Delete)
		v1.PATCH("/tags/:id/state", tagController.Update)

		//demo示例
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
