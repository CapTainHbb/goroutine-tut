package main

import (
	"fmt"
	"sync"
	"time"
)

func stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1_000_000; i++ {
		mutex.Lock()
		*money += 100
		mutex.Unlock()
	}
	fmt.Println("stingy done")
}

func spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1_000_000; i++ {
		mutex.Lock()
		*money -= 100
		mutex.Unlock()
	}
	fmt.Println("spendy done")
}

func main() {
	money := 100
	mutex := sync.Mutex{}
	go spendy(&money, &mutex)
	go stingy(&money, &mutex)

	time.Sleep(time.Second * 2)
	mutex.Lock()
	fmt.Println("money in bank: ", money)
	mutex.Unlock()
}