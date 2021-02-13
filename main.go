package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	apiKey := loadDotenvVariable("API_KEY")
	query := r.URL.Query().Get("query")

	str := fmt.Sprintln("https://core.ac.uk:443/api-v2/articles/search/?similar=false")

	u, _ := url.Parse(str)

	q := u.Query()
	q.Set("apiKey", apiKey)
	q.Set("query", query)
	u.RawQuery = q.Encode()

	url := u.String()

	response, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	// defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &sr)

	fmt.Println(response.Status)
	fmt.Println(sr.TotalHits)
	fmt.Println(sr.Status)

	if len(sr.Data) == 0 {
		log.Fatalln("")
	}

	fmt.Println(sr.Data[0].Title)
	fmt.Println(sr.Data[0].Description)
}

func loadDotenvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Dotenv is unreachable", err)
	}

	return os.Getenv(key)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/articles/search/{query}", getArticles).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
