package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func printFileContent(fileName string, textToFind string) {
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
	if len(os.Args) < 3 {
		fmt.Println("insufficient arguments!")
		return
	}

	args := os.Args[2:]
	textToFind := os.Args[1]

	for _, arg := range args {
		go printFileContent(arg, textToFind)
	}

	time.Sleep(5 * time.Second)
}