package server

import (
	"fmt"
	"log"
	"main/helpers"
	"main/routes"
	"main/structs"
	"net/http"
)

var Artists []structs.Artist

func Server() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := helpers.LoadDataFromAPI(&Artists)
		if err != nil {
			routes.RenderErrorPage(w, err, http.StatusInternalServerError)
			return
		}
		switch r.URL.Path {
		case "/", "/home":
			routes.RenderHomePage(w, r, Artists)
		case "/artist":
			routes.RenderArtistPage(w, r)
		default:
			routes.RenderNotFoundPage(w, r)
		}
	})

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path[7:]
		fullPath := "static/" + filePath
		http.ServeFile(w, r, fullPath)
	})

	errh := helpers.CheckTemplates()
	if errh != nil {
		return
	}

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
