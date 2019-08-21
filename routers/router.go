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

	// about
	r.GET("/about", v1.About)

	apiv1 := r.Group("/api/v1")
	{
		// blog list page
		apiv1.GET("/blogs", v1.BlogList)

		// blog categories page
		apiv1.GET("/categories", v1.Categories)

		// blog detail
		apiv1.GET("/detail", v1.Detail)
	}
	return r
}
