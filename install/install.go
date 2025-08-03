package install

import (
	"PackageManager/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
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
				fmt.Println("trying to install the sofware again")
		}




	}
}






func isLocalRepoUpToDate (repoPath string) bool{


	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("Could not open repo: %v", err)
	}

	/*
	// get local HEAD
	headRef, err := repo.Head()
	if err != nil {
		log.Fatalf("Could not get HEAD: %v", err)
	}
	localHash := headRef.Hash()
	fmt.Println("Local HEAD:", localHash) */

	// fetch latest from remote (but don't update local branches yet)
	start := time.Now()
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Depth: 1,
		Progress:   os.Stdout,
		//Progress: os.Stdout,
		// set Force: true if you want to always fetch (ignores "already up to date")
		//Force: true,
	})
	fmt.Println("Fetch took:", time.Since(start))

	if err != nil && err != git.NoErrAlreadyUpToDate {

		log.Fatalf("Fetch failed: %v", err)
	}



	if err == git.NoErrAlreadyUpToDate { // localHash == remoteHas ? 
		fmt.Println("✅ Remote is already up to date.")
		return true
	
	} else {
		updateRepoIfNeeded(repo)
		return false
	}






	/*
	// get remote branch reference //updating the local ref 
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
	}*/




}





func updateRepoIfNeeded(repo *git.Repository) {
	fmt.Println("📦 Updates found. Fast-forwarding local branch...")

	// Get remote main ref (origin/main)
	remoteRef, err := repo.Reference(plumbing.ReferenceName("refs/remotes/origin/main"), true)
	if err != nil {
		log.Fatalf("Could not get origin/main: %v", err)
	}

	// Update local main branch ref
	err = repo.Storer.SetReference(plumbing.NewHashReference(
		plumbing.ReferenceName("refs/heads/main"),
		remoteRef.Hash(),
	))
	if err != nil {
		log.Fatalf("Failed to fast-forward local main: %v", err)
	}

	// Reset working tree (i.e., "pull" the update)
	wt, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Could not get working tree: %v", err)
	}

	err = wt.Reset(&git.ResetOptions{
		Mode:   git.HardReset,
		Commit: remoteRef.Hash(),
	})
	if err != nil {
		log.Fatalf("Failed to reset working tree: %v", err)
	}

	fmt.Println("✅ Local main and working tree are now up to date with origin/main.")
}

