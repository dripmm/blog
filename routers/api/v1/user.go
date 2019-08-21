package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.html", nil)
}
