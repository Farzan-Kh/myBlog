package utils

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/about.html", "templates/home.html", "templates/post.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p any) {
	InfoLogger.Printf("Rendering template: %s", tmpl)
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		ErrorLogger.Printf("Error rendering template %s: %v", tmpl, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
