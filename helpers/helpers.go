package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// FetchJSON performs an HTTP GET request to the provided URL,
func FetchJSON(url string, target any) error {
	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("failed to fetch data from :" + url)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("failed to read response body from :" + url)
	}

	// Unmarshal JSON into the target
	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.New("failed to decode JSON from :" + url)
	}

	return nil
}

// LaunchBrowser tries to open the given URL in Google Chrome using the system's default command execution.
func LaunchBrowser(url string) error {
	err := exec.Command("google-chrome", url).Start()
	if err != nil {
		return fmt.Errorf("error opening browser:%w", err)
	}
	return nil
}

// checkTemplates verifies that all required HTML template files exist
// and are accessible by the server.
func CheckTemplates() error {
	templates := []string{"templates/error.html", "templates/index.html", "templates/404.html", "templates/artist.html"}
	for _, template := range templates {
		if _, err := os.Stat(template); os.IsNotExist(err) {
			return errors.New("template file " + template + " does not exist")
		}
	}
	return nil
}
