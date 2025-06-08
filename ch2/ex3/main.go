package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func searchForStringInFile(fileName string, textToFind string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
		return
	}

	if strings.Contains(string(file), textToFind) {
		fmt.Printf("found %s! %s\n", textToFind, fileName)
	}

}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("incorrect arguments!")
		return
	}

	dirToSearch := os.Args[2]
	textToFind := os.Args[1]

	files, err := os.ReadDir(dirToSearch) 
	if err != nil {
		fmt.Printf("error in reading directory %v\n", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fullpath := filepath.Join(dirToSearch, file.Name())
		abspath, err := filepath.Abs(fullpath)
		if err != nil {
			fmt.Printf("error getting abs path %v\n", err)
			continue
		}

		go searchForStringInFile(abspath, textToFind)
	}

	time.Sleep(5 * time.Second)
}