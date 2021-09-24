package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++{
		go func(numb int){
			fmt.Print("From main: ")
			fmt.Println(numb)
		}(i)
	}

	for i := 0; i < 10; i++{
		go doStuff(i)
	}

	time.Sleep(15 * time.Millisecond)
}

func doStuff(numb int)  {
	fmt.Print("From doStuff: ")
	fmt.Println(numb)
}