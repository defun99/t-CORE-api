package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("articles/search/", GetArticles)
	}
	return r
}

func main() {
	r := SetupRouter()
	r.Run()
}
