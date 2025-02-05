package main

import "net/http"

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	infoLogger.Println("Handling /about")
	renderTemplate(w, "about", nil)
}
