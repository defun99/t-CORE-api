package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type SearchResult struct {
	Status    string    `json:"status"`
	TotalHits int       `json:"totalHits"`
	Data      []Article `json:"data"`
}

type Article struct {
	Id          string   `json:"id"`
	Authors     []string `json:"authors"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Fulltext    string   `json:"fulltext"`
}

var articles []Article

func getArticles(w http.ResponseWriter, r *http.Request) {
	var sr SearchResult
	response, err := http.Get("https://core.ac.uk:443/api-v2/articles/search/cognitive?page=1&pageSize=10&metadata=true&fulltext=false&citations=false&similar=false&duplicate=false&urls=false&faithfulMetadata=false&apiKey=qdv5HZfMP9OY0pFb34BI1jQLNxD2kuiJ")

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &sr)

	fmt.Println(response.Status)
	fmt.Println(sr.TotalHits)
	fmt.Println(sr.Status)

	fmt.Println(sr.Data[0].Title)
	fmt.Println(sr.Data[0].Description)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/articles/search/{query}", getArticles).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
