package main

/*
special thanks to patrickfav for developing the open source apk signer and the aligner,
without him this project would have been impossible since google has removed any and all copies of their own apk signer

*/

import (
	"fmt"
	"os/exec"

	"android-apk-suite/cmd"
)

var packagemanadgers = [...]string{"wget", "curl", "javac"}

func main() {
	if !PMcheck() {
		fmt.Println("you dont have wget or curl or javac install them so you can utilize this program")
		return
	}
	//tool_downloader("36")
	cmd.Execute()
	//tool_downloader("36")

}

func PMcheck() bool {
	var check_downloaders = false
	for _, i := range packagemanadgers {
		_, err := exec.Command("which", string(i)).Output()
		if err != nil {
			continue
		}
		check_downloaders = true
	}
	return check_downloaders
}
