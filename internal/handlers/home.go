package handlers

import (
	"net/http"

	"myBlog/internal/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.InfoLogger.Println("Handing /home")
	pi_mu.Lock()
	defer pi_mu.Unlock()
	utils.RenderTemplate(w, "home", posts)
}
