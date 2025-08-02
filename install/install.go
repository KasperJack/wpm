package install
import (
	"PackageManager/config"
	"os"
	"path/filepath"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"fmt"
	"log"
)

func Install(software string) {

	expectedFile := software + ".json"
	targetPath := filepath.Join(config.LOCAL_MAIN_BUCKET, expectedFile)


	_, err := os.Stat(targetPath)

	if err == nil {
		//parce
		//get instaltion type
		//install

	}else{
		 // ---------> check if repo updated 
		if isLocalRepoUpToDate(config.LOCAL_MAIN_REPO){
			fmt.Println("sfoware does not exist")
		}else{

				//update main repo from github
				//check again
		}




	}
}






func isLocalRepoUpToDate (repoPath string) bool{


	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("Could not open repo: %v", err)
	}

	// get local HEAD
	headRef, err := repo.Head()
	if err != nil {
		log.Fatalf("Could not get HEAD: %v", err)
	}
	localHash := headRef.Hash()
	fmt.Println("Local HEAD:", localHash)

	// fetch latest from remote (but don't update local branches yet)
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		//Progress: os.Stdout,
		// set Force: true if you want to always fetch (ignores "already up to date")
		//Force: true,
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Fatalf("Fetch failed: %v", err)
	}

	// get remote branch reference
	remoteRef, err := repo.Reference(plumbing.ReferenceName("refs/remotes/origin/main"), true)
	if err != nil {
		log.Fatalf("Could not get remote ref: %v", err)
	}
	remoteHash := remoteRef.Hash()
	fmt.Println("Remote HEAD:", remoteHash)

	if localHash == remoteHash {
		fmt.Println("✅ Local repo is up to date with origin/main.")
		return true
	} else {
		fmt.Println("❌ Local repo is NOT up to date.")
		return false
	}




}