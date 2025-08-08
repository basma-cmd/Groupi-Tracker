package server

import (
	"fmt"
	"log"
	"main/helpers"
	"main/routes"
	"main/structs"
	"net/http"
	//"net/url"
	"os"
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
		info, err := os.Stat(r.URL.Path[1:])
		if err != nil || info.IsDir() {
			routes.RenderErrorPage(w, err, http.StatusForbidden)
			return
		}
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	errh := helpers.CheckTemplates()
	if errh != nil {
		return
	}
	helpers.LaunchBrowser("http://localhost:8080")

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
