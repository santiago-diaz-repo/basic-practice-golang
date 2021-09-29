package main

import "fmt"

func multiply(n1 int, n2 int, c chan int ){
	c <- n1 * n2
}

func main() {
	ch := make(chan int)
	go multiply(1,2, ch)
	go multiply(3,4, ch)

	n1 := <- ch
	n2 := <- ch

	fmt.Println(n1 * n2)
}
