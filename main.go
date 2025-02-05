package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

var templates = template.Must(template.ParseFiles("templates/about.html", "templates/home.html", "templates/post.html"))

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debug       bool
)

func renderTemplate(w http.ResponseWriter, tmpl string, p any) {
	infoLogger.Printf("Rendering template: %s", tmpl)
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		errorLogger.Printf("Error rendering template %s: %v", tmpl, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func timer(interval time.Duration, f func()) {
	infoLogger.Printf("Starting timer with interval: %v", interval)
	f()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		f()
	}

}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Could not load .env vars")
	}

	debug = os.Getenv("DEBUG") == "TRUE"

	logDir := "logs"
	logFile := filepath.Join(logDir, "logs.txt")

	// Create the directory if it doesn't exist
	if err = os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}

	infoLogger = log.New(file, "INFO: ", log.LUTC|log.Ltime|log.Ldate|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.LUTC|log.Ltime|log.Ldate|log.Lshortfile)

	index = make(map[string]*Post)
	unverified_emails = NewBiMap()

	go timer(1*time.Hour, FetchPosts)

	//Notif any, un-notified post
}

func main() {
	infoLogger.Println("Starting the application...")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/{$}", http.RedirectHandler("/home", http.StatusFound).ServeHTTP)

	http.HandleFunc("/posts/", postsHandler)

	http.HandleFunc("/newsletterReg", handleNewsletterReg)
	http.HandleFunc("/verifyEmail/", handleVerification)

	http.HandleFunc("/rss.xml", rssHandler)

	http.HandleFunc("/about/{$}", aboutHandler)

	http.HandleFunc("/search/", searchHandler)

	infoLogger.Println("Starting the web server...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
