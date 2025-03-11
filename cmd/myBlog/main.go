package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"myBlog/internal/config"
	"myBlog/internal/handlers"
	"myBlog/internal/utils"
)

var Debug bool

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Could not load .env vars")
	}
	config.Debug = os.Getenv("DEBUG") == "TRUE"

	utils.InitLogger(utils.InfoLogger, utils.ErrorLogger)

	go utils.Timer(1*time.Hour, handlers.FetchPosts) //Fetch posts every hour

	//Notif any, un-notified post
}

func main() {
	utils.InfoLogger.Println("Starting the application...")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/{$}", http.RedirectHandler("/home", http.StatusFound).ServeHTTP)

	http.HandleFunc("/posts/", handlers.PostsHandler)

	http.HandleFunc("/newsletterReg", handlers.HandleNewsletterReg)
	http.HandleFunc("/verifyEmail/", handlers.HandleVerification)

	http.HandleFunc("/rss.xml", handlers.RssHandler)

	http.HandleFunc("/about/{$}", handlers.AboutHandler)

	http.HandleFunc("/search/", handlers.SearchHandler)

	utils.InfoLogger.Println("Starting the web server...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
