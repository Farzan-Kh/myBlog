package models

import "html/template"

type Post struct {
	Title string        `json:"title"`
	Date  string        `json:"publishedAt"`
	UUID  string        `json:"uuid"`
	Body  template.HTML `json:"content"`
}
