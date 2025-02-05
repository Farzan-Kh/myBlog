package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"regexp"
	"sync"

	"github.com/gomarkdown/markdown"
	"github.com/valyala/fastjson"
)

type Post struct {
	Title string        `json:"title"`
	Date  string        `json:"publishedAt"`
	UUID  string        `json:"uuid"`
	Body  template.HTML `json:"content"`
}

type Response struct {
	Data []Post `json:"data"`
}

var posts []*Post
var index map[string]*Post
var pi_mu sync.Mutex

func FetchPosts() {
	if debug {
		FetchDummyPosts()
	} else {
		FetchStrapiPosts()
	}
}

func FetchDummyPosts() {
	infoLogger.Println("Fetching Dummy posts")
	file, err := os.Open("dummyData/dummyPosts.json")
	if err != nil {
		errorLogger.Printf("Error opening dummyPosts.json: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		errorLogger.Printf("Error reading dummyPosts.json: %v", err)
	}

	var p fastjson.Parser
	v, err := p.Parse(string(content))
	if err != nil {
		errorLogger.Printf("Error parsing dummy data: %v", err)
		return
	}

	pi_mu.Lock()
	defer pi_mu.Unlock()

	for _, vv := range v.GetArray("data") {
		uuid := string(vv.GetStringBytes("uuid"))
		p := &Post{
			Title: string(vv.GetStringBytes("title")),
			Date:  string(vv.GetStringBytes("publishedAt"))[:10],
			UUID:  uuid,
			Body:  template.HTML(markdown.ToHTML(vv.GetStringBytes("content"), nil, nil)),
		}
		index[uuid] = p
		posts = append(posts, p)
	}
	infoLogger.Println("Dummy posts processed and stored successfully")
}

func FetchStrapiPosts() {
	infoLogger.Println("Fetching posts from API")
	api_url := os.Getenv("API_ADDRESS") + "/blog-posts"

	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		errorLogger.Printf("Error creating API request: %v", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorLogger.Printf("Error making API request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorLogger.Printf("Error reading API response body: %v", err)
		return
	}

	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		errorLogger.Printf("Error parsing API response: %v", err)
		return
	}

	pi_mu.Lock()
	defer pi_mu.Unlock()

	for _, vv := range v.GetArray("data") {
		uuid := string(vv.GetStringBytes("uuid"))
		p := &Post{
			Title: string(vv.GetStringBytes("title")),
			Date:  string(vv.GetStringBytes("publishedAt"))[:10],
			UUID:  uuid,
			Body:  template.HTML(markdown.ToHTML(vv.GetStringBytes("content"), nil, nil)),
		}
		index[uuid] = p
		posts = append(posts, p)
	}
	infoLogger.Println("Posts fetched and stored successfully")
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	infoLogger.Println("Handling /posts/")
	uuid := r.URL.Path[len("/posts/"):]

	re := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	if !re.MatchString(uuid) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid UUID")
		errorLogger.Printf("Invalid UUID: %s", uuid)
		return
	}

	pi_mu.Lock()
	defer pi_mu.Unlock()
	if p, ok := index[uuid]; ok {
		renderTemplate(w, "post", p)
	} else {
		w.WriteHeader(http.StatusNotFound)
		errorLogger.Printf("Post not found: %s", uuid)
	}
}
