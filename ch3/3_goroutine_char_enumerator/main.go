package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"


func countLetters(url string, frequency []int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("server returning status code: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Printf("Completed %v\n", url)
}

func main() {
	frequency := make([]int, 26)

	for i := 1000; i <= 1005; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency)
	}


	time.Sleep(5 * time.Second)

	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
}