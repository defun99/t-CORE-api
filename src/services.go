package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"net/url"
)

var myEnv map[string]string
var baseUrl string
var apiKey string
var client *http.Client

func init() {
	myEnv, err := godotenv.Read()
	baseUrl = myEnv["BASE_URL"]
	apiKey = myEnv["API_KEY"]

	if err != nil {
		fmt.Println(".env file is not provided")
	}

	client = &http.Client{}
}

func RetrieveArticleSearchResult(searchQuery string, page string, pageSize string) (SearchResult, error) {

	var searchResponse SearchResult
	var articles []Article

	base, err := url.Parse(baseUrl)
	if err != nil || base == nil {
		fmt.Println("Invalid base baseUrl")
	}

	// Build baseUrl
	baseUrl += fmt.Sprintf("articles/search/%s", searchQuery)
	base.ForceQuery = true

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return SearchResult{}, err
	}

	q := req.URL.Query()
	q.Add("apiKey", apiKey)
	q.Add("page", page)
	q.Add("pageSize", pageSize)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL)

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Content-Length", "6907")
	resp, err := client.Do(req)

	if err != nil {
		return SearchResult{}, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error: ", err)
	}
	err = json.Unmarshal(content, &searchResponse)
	resp.Body.Close()

	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(searchResponse)
	fmt.Println(articles)

	return searchResponse, nil
}
