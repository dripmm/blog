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

	// blog list
	r.GET("/blogs", v1.BlogList)

	// categories
	r.GET("/categories", v1.Categories)

	// blog dateil
	r.GET("/detail", v1.Detail)

	// user sing in
	r.GET("/sign", v1.SignIn)

	// article add
	r.GET("/article", v1.Article)
	//apiv1 := r.Group("/api/v1")
	//{
	//}
	return r
}
