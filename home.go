package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	infoLogger.Println("Handing /home")
	pi_mu.Lock()
	defer pi_mu.Unlock()
	renderTemplate(w, "home", posts)
}
