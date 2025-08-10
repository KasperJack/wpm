package install



import (
	"encoding/json"
	"fmt"
	"os"
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
	//fmt.Println(m.)


}