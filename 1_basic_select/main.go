package main

import (
	"fmt"
	"time"
)


func createProducer(msgToProduce string, timeout time.Duration) (<-chan string) {
	outChan := make(chan string)
	go func(){
		for {
			time.Sleep(timeout)
			outChan <- msgToProduce
		}
	} ()
	return	outChan
}

func main() {
	msgChanA := createProducer("tick", time.Second * 3)
	msgChanB := createProducer("tock", time.Second * 1)

	for {
		select {
		case msgA := <- msgChanA:
			fmt.Println(msgA)
		case msgB := <- msgChanB:
			fmt.Println(msgB)
		}
	}
}