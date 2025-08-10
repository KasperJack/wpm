package main

import (
	"fmt"
	//"github.com/go-git/go-git/v6"
	"PackageManager/install"
	"os"
)

func main(){
	
	if (len(os.Args) ==1){
		fmt.Println("add agrs")
		os.Exit(0)
	}	


	if (os.Args[1] == "install") {
		install.Install("test")


	}else if(os.Args[1] == "update"){

		install.UpdateLocalRepo()

	}
	

}