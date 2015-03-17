package xmd

import (
"fmt"
)

const (
	gitConfigPath = ".git/config"
	commitCommandPath = ".git/hooks/xmd-commit"
	
)

var (
	commitCommand = []byte("echo 'xmd was here'\n")
)

func SetupGit(){
	configExists, err := Exists(gitConfigPath)
	Check(err)
	if configExists {
		//write commit hook
		err := WriteBytesToFile(commitCommandPath,commitCommand)
		if err == nil {
			fmt.Println("written commit hook")
		}
		Check(err)

		//TODO write pull/clone hook

	} else {
		fmt.Println("its not a valid git repo")
	}
}

