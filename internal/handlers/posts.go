package handlers

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

	"myBlog/internal/config"
	"myBlog/internal/models"
	"myBlog/internal/utils"
)

var posts []*models.Post
var index map[string]*models.Post = make(map[string]*models.Post)
var pi_mu sync.Mutex

func FetchPosts() {
	if config.Debug {
		FetchDummyPosts()
	} else {
		FetchStrapiPosts()
	}
}

func FetchDummyPosts() {
	utils.InfoLogger.Println("Fetching Dummy posts")
	file, err := os.Open("dummyData/dummyPosts.json")
	if err != nil {
		utils.ErrorLogger.Printf("Error opening dummyPosts.json: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		utils.ErrorLogger.Printf("Error reading dummyPosts.json: %v", err)
	}

	var p fastjson.Parser
	v, err := p.Parse(string(content))
	if err != nil {
		utils.ErrorLogger.Printf("Error parsing dummy data: %v", err)
		return
	}

	pi_mu.Lock()
	defer pi_mu.Unlock()

	for _, vv := range v.GetArray("data") {
		uuid := string(vv.GetStringBytes("uuid"))
		p := &models.Post{
			Title: string(vv.GetStringBytes("title")),
			Date:  string(vv.GetStringBytes("publishedAt"))[:10],
			UUID:  uuid,
			Body:  template.HTML(markdown.ToHTML(vv.GetStringBytes("content"), nil, nil)),
		}
		index[uuid] = p
		posts = append(posts, p)
	}
	utils.InfoLogger.Println("Dummy posts processed and stored successfully")
}

func FetchStrapiPosts() {
	utils.InfoLogger.Println("Fetching posts from API")
	api_url := os.Getenv("API_ADDRESS") + "/blog-posts"

	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		utils.ErrorLogger.Printf("Error creating API request: %v", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.ErrorLogger.Printf("Error making API request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.ErrorLogger.Printf("Error reading API response body: %v", err)
		return
	}

	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		utils.ErrorLogger.Printf("Error parsing API response: %v", err)
		return
	}

	pi_mu.Lock()
	defer pi_mu.Unlock()

	for _, vv := range v.GetArray("data") {
		uuid := string(vv.GetStringBytes("uuid"))
		p := &models.Post{
			Title: string(vv.GetStringBytes("title")),
			Date:  string(vv.GetStringBytes("publishedAt"))[:10],
			UUID:  uuid,
			Body:  template.HTML(markdown.ToHTML(vv.GetStringBytes("content"), nil, nil)),
		}
		index[uuid] = p
		posts = append(posts, p)
	}
	utils.InfoLogger.Println("Posts fetched and stored successfully")
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println("Handling /posts/")
	uuid := r.URL.Path[len("/posts/"):]

	re := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	if !re.MatchString(uuid) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid UUID")
		utils.ErrorLogger.Printf("Invalid UUID: %s", uuid)
		return
	}

	pi_mu.Lock()
	defer pi_mu.Unlock()
	if p, ok := index[uuid]; ok {
		utils.RenderTemplate(w, "post", p)
	} else {
		w.WriteHeader(http.StatusNotFound)
		utils.ErrorLogger.Printf("Post not found: %s", uuid)
	}
}
