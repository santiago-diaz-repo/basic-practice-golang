package main

import (
	"fmt"
	"time"
)

func sum(num1 int, num2 int, ch chan int) {
	ch <- num1 + num2
}

func multiply(num1 int, num2 int, ch chan int) {
	ch <- num1 * num2
	//<- ch // empty this channel, or consumes whatever piece of data it has
}

func exec(sum , multiply chan int){
	select {
	case  b := <-sum:
		fmt.Printf("sum is: %s \n",b)
	case a := <- multiply:
		fmt.Printf("sum is: %s \n",a)
	default:
		fmt.Printf("Nothing here \n")
	}
}

func main() {
	chSum := make(chan int)
	chMultiply := make(chan int)

	for i := 0; i < 2; i++{
		go sum(i, i+1, chSum)
		go multiply(i, i+1, chMultiply)
	}
	for i := 0; i < 5; i++ {
		go exec(chSum, chMultiply)
	}
	time.Sleep(20 * time.Millisecond)
}
