package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IGetArticles interface {
	GetArticles(c *gin.Context)
}

func GetArticles(c *gin.Context, searchQuery string, page string, pageSize string) {

	articles, err := RetrieveArticleSearchResult(searchQuery, page, pageSize)

	if err != nil {
		fmt.Println(err)
	}
	if articles == nil || len(articles) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, articles)
		return
	}
}
