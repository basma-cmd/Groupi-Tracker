package main

import (
	"main/helpers"
)

// LoadDataFromAPI fetches artist, location, date, and relation data
// from the Groupie Tracker API and stores them in global variables.
func LoadDataFromAPI() error {
	// Fetch artist data
	err := helpers.FetchJSON("https://groupietrackers.herokuapp.com/api/artists", &Artists)
	if err != nil {
		return err
	}
	return nil
}
