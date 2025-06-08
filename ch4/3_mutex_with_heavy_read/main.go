package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)


func matchRecorder(matchEvents *[]string, mutex *sync.Mutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		*matchEvents = append(*matchEvents, "match event " + strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(time.Millisecond * 200)
		fmt.Println("appended match event")
	}
}

func copyAllEvents(matchEvents *[]string) []string {
	return (*matchEvents)[:]
}

func clientHandler(mEvents *[]string, mutex *sync.Mutex, st time.Time) {
	for i := 0; i < 100; i++ {
		mutex.Lock()
		allEvents := copyAllEvents(mEvents)
		mutex.Unlock()

		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), " events copied in ", timeTaken)
	}
}

func main() {
	mutex := sync.Mutex{}
	matchEvents := make([]string, 0, 10_000)
	for i := 0; i < 10_000; i++ {
		matchEvents = append(matchEvents, "match event")
	}
	go matchRecorder(&matchEvents, &mutex)
	start := time.Now()

	for j := 0; j < 5_000; j++ {
		go clientHandler(&matchEvents, &mutex, start)
	}
	time.Sleep(100 * time.Second)
}