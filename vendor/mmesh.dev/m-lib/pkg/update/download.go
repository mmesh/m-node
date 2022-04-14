package update

import (
	"io"
	"net/http"
	"os"
)

// downloadFile download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filename string) error {
	//filename := path.Base(url)

	url := getURL(filename)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the body to file
	_, err = io.Copy(f, resp.Body)

	return err
}
