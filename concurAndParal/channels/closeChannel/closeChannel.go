package main

import (
	"fmt"
	"time"
)

func receive(ch chan int){

	if data, ok := <- ch; ok{
		fmt.Printf("Receiving: %n ", data)
	}

}

func send(num int, ch chan<- int){
	fmt.Println("sending")

	ch <- num
	close(ch)
}

func main() {
	ch := make(chan int)
	go send(4, ch)

	go receive(ch)
	go receive(ch)
	time.Sleep(6 * time.Millisecond)

}

