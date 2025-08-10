package install

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"PackageManager/config"
	"io"
)

type Manifest struct {
	Version     string `json:"version"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	License     string `json:"license"`
	Notes       string `json:"notes"`
	Arch        struct {
		X64 struct {
			URL  string `json:"url"`
			Hash string `json:"hash"`
		} `json:"64bit"`
	} `json:"arch"`
	Bin string `json:"bin"`
}



func parseJson(filePath string){

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}


	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		panic(err)
	}

	// Example usage
	fmt.Println("Version:", m.Version)
	fmt.Println("Download URL:", m.Arch.X64.URL)


	fileName := "btop4win-LHM-x64.zip"

	targetPath := filepath.Join(config.APPS_DIR, fileName)


	// Download
	err = downloadFile(m.Arch.X64.URL, targetPath)

	if err != nil {
		fmt.Println("Download error:", err)
		return
	}

	fmt.Println("File downloaded:", targetPath)










}






// downloadFile downloads from URL and saves to a file
func downloadFile(url, filepath string) error {
	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy data from response to file
	_, err = io.Copy(out, resp.Body)
	return err
}