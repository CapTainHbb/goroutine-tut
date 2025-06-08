package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"


func countLetters(url string, frequency []int, mutex *sync.Mutex) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("server returning status code: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	mutex.Lock()
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Printf("Completed %v\n", url)
	mutex.Unlock()
}

func main() {
	mutex := sync.Mutex{}
	frequency := make([]int, 26)

	for i := 1000; i <= 1010; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mutex)
	}
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		if mutex.TryLock() {
			for i, c := range allLetters {
				fmt.Printf("%c-%d ", c, frequency[i])
			}
			mutex.Unlock() // very important
		} else {
			fmt.Println("mutex already being used")
		}
	}

	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
}