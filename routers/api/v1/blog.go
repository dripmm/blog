package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BlogList(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/index.html", nil)
}

func Categories(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/categories.html", nil)
}

func Detail(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/detail.html", nil)
}

// admin article add
func Article(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/article_add.html", nil)
}
