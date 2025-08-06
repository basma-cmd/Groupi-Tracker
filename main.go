package main

import (
	"fmt"
	"log"
	"net/http"

	"main/helpers"
	"main/routes"
	"main/structs"
)

// Global variables to hold data fetched from the API
var Artists []structs.Artist

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Load data from external APIs
		err := LoadDataFromAPI()
		if err != nil {
			// If fetching data fails, show error page
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

	// Serve static CSS files from the /style/ directory
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path[7:]
		fullPath := "static/" + filePath // Build full path to file
		http.ServeFile(w, r, fullPath)   // Serve the file
	})
	// Automatically open browser to the home page
	helpers.LaunchBrowser("http://localhost:8080")
	err := helpers.CheckTemplates()
	if err != nil {
		return 
	}
	
	fmt.Println("Server running at http://localhost:8080")

	// Start the HTTP server and handle fatal errors
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
