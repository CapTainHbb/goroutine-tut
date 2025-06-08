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

	for {
		select {
		case msgA := <- msgChanA:
			fmt.Println(msgA)
			return
		default:
			fmt.Println("default case")
			time.Sleep(time.Second * 1)
		}
	}
}