package controllers

import (
	"fmt"
	models "main/app/models"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var myEnv map[string]string
var baseUrl string
var apiKey string

func init() {
	myEnv, err := godotenv.Read()
	baseUrl = myEnv["BASE_URL"]
	apiKey = myEnv["API_KEY"]

	if err != nil {
		fmt.Println(".env file is not provided")
	}
}

func GetArticles(c *gin.Context) {
	var articles []models.Article

	base, err := url.Parse(baseUrl)
	if err != nil {
		return
	}

	// Eject parameters from client request
	searchQuery := c.Request.URL.Query().Get("query")
	page := c.Request.URL.Query().Get("page")
	pageSize := c.Request.URL.Query().Get("pageSize")

	// Build url
	base.Path += searchQuery
	base.ForceQuery = true
	params := url.Values{}
	params.Add("apiKey", apiKey)
	params.Add("page", page)
	params.Add("pageSize", pageSize)
	base.RawQuery = params.Encode()

	c.BindJSON(&articles)

	url := myEnv["BASE_URL"]
	apiKey := myEnv["API_KEY"]

	fmt.Println(apiKey)

	req, err := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("apKey", apiKey)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	// TODO: Use model to parse response and save to articles
	// TODO: Send response to client

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		c.JSON(http.StatusOK, articles)
		return
	}
}
