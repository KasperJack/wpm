package install
import (
	"PackageManager/config"
	"os"
	"path/filepath"
	//"fmt"
)

func Install(software string) {

	expectedFile := software + ".json"
	targetPath := filepath.Join(config.LOCAL_MAIN, expectedFile)


	_, err := os.Stat(targetPath)

	if err == nil {
		//parce
		//install

	}else{
		// check if repo updated 
		//update main repo from github
		//check again
	}





}