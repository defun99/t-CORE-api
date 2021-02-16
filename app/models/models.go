package models

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
