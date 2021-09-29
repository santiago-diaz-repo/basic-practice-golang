package main

import (
	"fmt"
	"time"
)

func receive(ch <-chan int){
	num := <- ch
	fmt.Printf("Receiving: %n ", num)
}

func send(num int, ch chan<- int){
	fmt.Println("sending")
	ch <- num
}

func main() {
	ch := make(chan int)
	go send(4, ch)
	go receive(ch)
	time.Sleep(6 * time.Millisecond)
}
