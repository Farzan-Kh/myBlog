package handlers

import (
	"net/http"

	"myBlog/internal/utils"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println("Handling /about")
	utils.RenderTemplate(w, "about", nil)
}
