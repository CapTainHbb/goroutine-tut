package main

import (
	"fmt"
	"os"
	"time"
)

func printFileContent(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
		return
	}

	fmt.Println(string(file))

}

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		go printFileContent(arg)
	}

	time.Sleep(5 * time.Second)
}