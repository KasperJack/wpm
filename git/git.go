package git

import (
	"PackageManager/config"
	"os/exec"
	"os"
	"log"
)


func Clone(targeRepo string, dir string) {

	//fetch
	clone := exec.Command(config.GIT,"clone", targeRepo)
	clone.Dir = dir
	
    clone.Stdout = os.Stdout
    clone.Stderr = os.Stderr
    if err := clone.Run(); err != nil {
        log.Fatal(err)
        return
    }


}
func FetchDry (dir string) []byte{


	fetchDry := exec.Command(config.GIT,"fetch", "--dry-run", "origin")
	fetchDry.Dir = dir
	
	output, err := fetchDry.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }

	return output
}



func Fetch(dir string) {

		//fetch
	fetch := exec.Command(config.GIT,"fetch", "origin")
	fetch.Dir = dir
	
    fetch.Stdout = os.Stdout
    fetch.Stderr = os.Stderr
    if err := fetch.Run(); err != nil {
        log.Fatal(err)
        return
    }
}




func Pull(dir string){

	pull := exec.Command(config.GIT, "pull", "origin", "main")
    pull.Dir = dir
    pull.Stdout = os.Stdout
    pull.Stderr = os.Stderr

    if err := pull.Run(); err != nil {
        log.Fatal(err)
        return
    }

}