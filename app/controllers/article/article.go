package article

import "github.com/gin-gonic/gin"

type Article struct {
}

var ArticleController = &Article{}

func Get(c *gin.Context) {
	ArticleController.Get(c)
}

func List(c *gin.Context) {
	ArticleController.List(c)
}

func Create(c *gin.Context) {
	ArticleController.Create(c)
}

func Update(c *gin.Context) {
	ArticleController.Update(c)
}

func Delete(c *gin.Context) {
	ArticleController.Delete(c)
}
func (a *Article) Get(c *gin.Context) {}
func (a *Article) List(c *gin.Context) {}
func (a *Article) Create(c *gin.Context) {}
func (a *Article) Update(c *gin.Context) {}
func (a *Article) Delete(c *gin.Context) {}