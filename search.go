package main

import (
	"net/http"
	"net/url"
	"strings"
)

func search(search_term string) []Post {
	infoLogger.Printf("Searching for term: %s", search_term)
	search_results := make([]Post, 0, 10)

	for _, value := range index {
		if strings.Contains(string(value.Body), search_term) || strings.Contains(string(value.Title), search_term) {
			search_results = append(search_results, *value)
		}
	}

	return search_results
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	infoLogger.Println("Handling /search/")
	search_term := r.URL.Path[len("/search/"):]
	search_term = url.QueryEscape(search_term)

	// search_results := search(search_term)
}
