package routers

import (
	controllers "main/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("articles/search/", controllers.GetArticles)
	}
	return r
}
