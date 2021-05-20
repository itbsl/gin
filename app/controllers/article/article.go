package article

import "github.com/gin-gonic/gin"

type Article struct {
}

var a = &Article{}

func Get(c *gin.Context) {
	a.Get(c)
}

func List(c *gin.Context) {
	a.List(c)
}

func Create(c *gin.Context) {
	a.Create(c)
}

func Update(c *gin.Context) {
	a.Update(c)
}

func Delete(c *gin.Context) {
	a.Delete(c)
}
func (a *Article) Get(c *gin.Context) {}
func (a *Article) List(c *gin.Context) {}
func (a *Article) Create(c *gin.Context) {}
func (a *Article) Update(c *gin.Context) {}
func (a *Article) Delete(c *gin.Context) {}