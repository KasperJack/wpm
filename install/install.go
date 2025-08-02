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
	targetPath := filepath.Join(config.LOCAL_MAIN, expectedFile)


	_, err := os.Stat(targetPath)

	if err == nil {
		//parce
		//get instaltion type
		//install

	}else{
		 // ---------> check if repo updated 
		//update main repo from github
		//check again

		repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("Could not open repo: %v", err)
	}

	// Get local HEAD
	headRef, err := repo.Head()
	if err != nil {
		log.Fatalf("Could not get HEAD: %v", err)
	}
	localHash := headRef.Hash()
	fmt.Println("Local HEAD:", localHash)

	// Fetch latest from remote (but don't update local branches yet)
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		// You can set `Progress: os.Stdout` to debug
		// set Force: true if you want to always fetch (ignores "already up to date")
		Force: true,
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Fatalf("Fetch failed: %v", err)
	}

	// Get remote branch reference
	remoteRef, err := repo.Reference(plumbing.ReferenceName("refs/remotes/origin/main"), true)
	if err != nil {
		log.Fatalf("Could not get remote ref: %v", err)
	}
	remoteHash := remoteRef.Hash()
	fmt.Println("Remote HEAD:", remoteHash)

	// Compare hashes
	if localHash == remoteHash {
		fmt.Println("✅ Local repo is up to date with origin/main.")
	} else {
		fmt.Println("❌ Local repo is NOT up to date.")
	}


	}


}


