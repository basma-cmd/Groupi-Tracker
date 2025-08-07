package helpers

import (
	"main/structs"
)

// LoadDataFromAPI fetches artist data from the Groupie Tracker API.
func LoadDataFromAPI(artists *[]structs.Artist) error {
	err := FetchJSON("https://groupietrackers.herokuapp.com/api/artists", artists)
	if err != nil {
		return err
	}
	return nil
}
