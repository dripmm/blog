package routers

import (
	"github.com/drip/hello/pkg/setting"
	v1 "github.com/drip/hello/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.LoadHTMLGlob("views/**/*")

	gin.SetMode(setting.ServerSetting.RunMode)

	// home page
	r.GET("/", v1.HomePage)
	//apiv1 := r.Group("/api/v1")
	//{
	//	apiv1.GET("/", v1.HomePage)
	//}
	return r
}
