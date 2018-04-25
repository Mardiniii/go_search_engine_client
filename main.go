package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Page struct to store in database
type Page struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// SearchResult struct to handle search queries
type SearchResult struct {
	Pages []Page `json:"pages"`
	Input string `json:"input"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/home.html")
	if err != nil {
		log.Print("Template parsing error: ", err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Print("Template executing error: ", err)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	searchInput := r.Form.Get("input")

	log.Print("Querying database for: ", searchInput)

	pages := SearchContent(searchInput)

	searchResult := SearchResult{
		Input: searchInput,
		Pages: pages,
	}

	jsonData, err := json.Marshal(searchResult)
	if err != nil {
		log.Print("JSON executing error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	NewElasticSearchClient()
	exists := ExistsIndex(indexName)

	if !exists {
		CreateIndex(indexName)
	}

	mux := mux.NewRouter()

	mux.HandleFunc("/", homeHandler).Methods("GET")
	mux.HandleFunc("/search", searchHandler).Methods("GET")
	http.ListenAndServe(":8080", mux)
}
