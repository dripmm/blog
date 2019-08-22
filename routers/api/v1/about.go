package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func About (c *gin.Context) {
	c.HTML(http.StatusOK, "front/about/index.html", nil)
}
