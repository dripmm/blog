package routers

import (
	"github.com/drip/hello/pkg/setting"
	"github.com/drip/hello/routers/admin"
	v1 "github.com/drip/hello/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.LoadHTMLGlob("views/**/**/*")

	gin.SetMode(setting.ServerSetting.RunMode)

	/* front pages */
	front := r.Group("/")
	{
		front.GET("/", v1.HomePage)
		front.GET("/about", v1.About)
		front.GET("/blogs", v1.BlogList)
		front.GET("/categories", v1.Categories)
		front.GET("/detail", v1.Detail)
	}
	backEnd := r.Group("/admin")
	{
		backEnd.GET("/", admin.SignIn)
		backEnd.GET("/article_add", admin.ArticleAdd)
	}
	//// user sing in
	//r.GET("/sign", admin.SignIn)
	//
	//// article add
	//r.GET("/article", admin.Article)
	//
	//// categoies add
	//r.GET("/categories_add", admin.CategoriesAdd)

	/* admin pages start */
	//apiv1 := r.Group("/api/v1")
	//{
	//}
	return r
}
