package handlers

import (
	"encoding/xml"
	"myBlog/internal/utils"
	"net/http"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

func RssHandler(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println("Handling /rss.xml")
	items := []Item{}
	pi_mu.Lock()
	defer pi_mu.Unlock()
	for _, post := range posts {
		items = append(items, Item{
			Title:       post.Title,
			Link:        "https://farzan.info/posts/" + post.UUID,
			Description: string(post.Body[:50]),
			PubDate:     post.Date,
			GUID:        "https://farzan.info/posts/" + post.UUID,
		})
	}

	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:       "Farzan's Blog",
			Link:        "https://farzan.info/posts",
			Description: "",
			Items:       items,
		},
	}

	w.Header().Set("Content-Type", "application/xml")
	err := xml.NewEncoder(w).Encode(rss)
	if err != nil {
		utils.ErrorLogger.Printf("Error encoding RSS feed: %v", err)
	}
}
