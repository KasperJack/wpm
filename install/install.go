package install

import (
	"PackageManager/config"
	"PackageManager/git"
	"fmt"
	//"log"
	"os"
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


	output := git.FetchDry(config.LOCAL_MAIN_REPO)


	if (len(output) == 0){
		fmt.Println("local repo up to date")
		return true
	}else{

		fmt.Println(string(output))
		return false
	}



}





func updateLocalRepo() {

	//fetch
	git.Fetch(config.LOCAL_MAIN_REPO)


	//pull

	git.Pull(config.LOCAL_MAIN_REPO)


}

