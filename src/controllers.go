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

	searchResult, err := RetrieveArticleSearchResult(searchQuery, page, pageSize)

	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, searchResult)
		return
	}
}
