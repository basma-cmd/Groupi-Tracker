package routes

import (
	"errors"
	"net/http"
	"text/template"

	"main/helpers"
	"main/structs"
)

// RenderNotFoundPage handles 404 (Not Found) errors.
func RenderNotFoundPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	tmpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		// Fall back to a generic error message if template parsing fails
		RenderErrorPage(w, err, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		// Fall back to generic error if template rendering fails
		RenderErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}

// RenderHomePage handles the GET request to the root URL ("/").
func RenderHomePage(w http.ResponseWriter, r *http.Request, artists any) {
	if r.Method != "GET" {
		RenderErrorPage(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderErrorPage(w, errors.New("error loading template"), http.StatusInternalServerError)
		return
	}

	// Render the template with an empty data map (no dynamic content)
	err = tmpl.Execute(w, artists)
	if err != nil {
		RenderErrorPage(w, errors.New("error loading template"), http.StatusInternalServerError)
		return
	}
}

// RenderErrorPage renders an HTML error page using the "error.html" template.
func RenderErrorPage(w http.ResponseWriter, details error, statusCode int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Error loading error template", http.StatusInternalServerError)
		return
	}

	// set the status code only
	w.WriteHeader(statusCode)

	data := map[string]error{
		"Details": details,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering error template", http.StatusInternalServerError)
	}
}

// RenderArtistPage renders the details page for a specific artist.
// It extracts the artist's ID from the query parameters, fetches the
// corresponding artist, locations, dates, and relations, and passes them to the template.
func RenderArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		RenderErrorPage(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	var Artist structs.Artist
	err := helpers.FetchJSON("https://groupietrackers.herokuapp.com/api/artists/"+idStr, &Artist)
	if err != nil {
		RenderErrorPage(w, errors.New("error fetching locations"), http.StatusBadRequest)
		return
	}

	var Locations structs.Locations
	err = helpers.FetchJSON("https://groupietrackers.herokuapp.com/api/locations/"+idStr, &Locations)
	if err != nil {
		RenderErrorPage(w, errors.New("error fetching locations"), http.StatusBadRequest)
		return
	}

	var Dates structs.Dates
	err = helpers.FetchJSON("https://groupietrackers.herokuapp.com/api/dates/"+idStr, &Dates)
	if err != nil {
		RenderErrorPage(w, errors.New("error fetching dates"), http.StatusBadRequest)
		return
	}

	var Relations structs.Relations
	err = helpers.FetchJSON("https://groupietrackers.herokuapp.com/api/relation/"+idStr, &Relations)
	if err != nil {
		RenderErrorPage(w, errors.New("error fetching relations"), http.StatusBadRequest)
		return
	}

	// Combine all fetched data into a struct to send to the template
	details := structs.ArtistDetails{
		Artist:    Artist,
		Locations: Locations,
		Dates:     Dates,
		Relations: Relations,
	}

	// Load and execute the artist details template
	tmpl, _ := template.ParseFiles("templates/artist.html")
	err = tmpl.Execute(w, details)
	if err != nil {
		// If the template fails to render, return a 500 error
		RenderErrorPage(w, errors.New("failed to render artist page"), http.StatusInternalServerError)
		return
	}
}
