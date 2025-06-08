package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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
}

func main() {
	frequency := make([]int, 26)

	for i := 1000; i <= 1005; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		countLetters(url, frequency)
		fmt.Printf("Completed %v\n", url)
	}

	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
}