package main

import (
	"log"
	"net/http"

	"myBlog/internal/utils"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Could not load .env vars")
	}

	utils.InitLogger(utils.InfoLogger, utils.ErrorLogger)
}

func main() {
	utils.InfoLogger.Println("Starting the application...")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/home.html")
	})

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[
			{"date": "2025-01-10", "title": "WorstFit: Unveiling Hidden Transformers in Windows ANSI!"},
			{"date": "2024-08-09", "title": "Confusion Attacks: Exploiting Hidden Semantic Ambiguity in Apache HTTP Server!"}
		]`))
	})

	utils.InfoLogger.Println("Server running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
