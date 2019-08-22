package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// admin article add
func ArticleAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/blog/article_add.html", nil)
}

// categories add page
func CategoriesAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "xxx", nil)
}
