package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


var words = []string{"the", "car", "program", "this"}


func countWords(url string, frequency []int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("server returning status code: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	c := strings.ToLower(string(body))
	tokens := strings.Fields(c)

	for _, token := range tokens {
		for i, word := range words {
			if word == token {
				frequency[i] += 1
			}
		}
	}
}

func main() {
	frequency := make([]int, 4)

	for i := 1000; i <= 1005; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		countWords(url, frequency)
		fmt.Printf("Completed %v\n", url)
	}
	
	for i, c := range words {
		fmt.Printf("%v-%d ", c, frequency[i])
	}
}