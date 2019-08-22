package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/user/login.html", nil)
}
