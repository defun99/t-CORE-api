package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.Status(200)
		})
		v1.GET("articles/search/", func(c *gin.Context) {
			page := c.DefaultQuery("firstname", "1")
			pageSize := c.DefaultQuery("pageSize", "1")
			searchQuery := c.Query("query")

			GetArticles(c, searchQuery, page, pageSize)
		})
	}

	s := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
