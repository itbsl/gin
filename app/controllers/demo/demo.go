package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 静态页面功能示例
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "demo/index.html", "《小萝莉的猴神大叔》")
}

// GetParam 带参数的路由示例：路径中直接加上参数值 127.0.0.1:8080/v1/demo/GetQuery/itbsl
func GetParam(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(name)
	c.String(http.StatusOK, "GetParam %s", name)
}

// GetQuery 带参数的路由示例：路径中使用参数名 127.0.0.1:8080/v1/demo/GetQuery?name=itbsl
func GetQuery(c *gin.Context) {
	name := c.Query("name")
	fmt.Println(name)
	c.String(http.StatusOK, "GetQuery %s", name)
}

// Post 获取POST请求数据示例
func Post(c *gin.Context) {
	name := c.PostForm("username")
	age := c.PostForm("age")
	c.String(200,"hello,%s,年龄为:%s",name,age)
}

// Ping 返回json格式数据示例
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "demo.Pong",
	})
}
