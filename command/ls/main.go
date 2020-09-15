package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		// ls .
	}
	targetDirPath := os.Args[1]
	if dirList, err := ioutil.ReadDir(targetDirPath); err == nil {
		for _, dirInfo := range dirList {
			fmt.Println(dirInfo.Name())
		}
	} else {
		fmt.Println(err.Error())
	}

}
