package install

import (
	"PackageManager/config"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"


	//"time"


)

func Install(software string) {




	
	expectedFile := software + ".json"
	targetPath := filepath.Join(config.LOCAL_MAIN_BUCKET, expectedFile)


	_, err := os.Stat(targetPath)

	if err == nil {
		parseJson(targetPath)
		//get instaltion type
		//install

	}/*else{
		 // ---------> check if repo updated 
		if isLocalRepoUpToDate(){
			fmt.Println("sfoware does not exist")
		}else{

				//update main repo from github
				UpdateLocalRepo()
				
				//check again
				fmt.Println("trying to install the sofware again")
		}




	}*/
}






func isLocalRepoUpToDate () bool{


	fetchDry := exec.Command(config.GIT,"fetch", "--dry-run", "origin")
	fetchDry.Dir = config.LOCAL_MAIN_REPO
	
	output, err := fetchDry.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }

	if (len(output) == 0){
		fmt.Println("local repo up to date")
		return true
	}else{

		fmt.Println(string(output))
		return false
	}




}





func UpdateLocalRepo() {

	//fetch
	fetch := exec.Command(config.GIT,"fetch", "origin")
	fetch.Dir = config.LOCAL_MAIN_REPO
	
    fetch.Stdout = os.Stdout
    fetch.Stderr = os.Stderr
    if err := fetch.Run(); err != nil {
        fmt.Println("Fetch failed:", err)
        return
    }


	//pull


	pull := exec.Command(config.GIT, "pull", "origin", "main")
    pull.Dir = config.LOCAL_MAIN_REPO
    pull.Stdout = os.Stdout
    pull.Stderr = os.Stderr

    if err := pull.Run(); err != nil {
        fmt.Println("Pull failed:", err)
        return
    }





}

