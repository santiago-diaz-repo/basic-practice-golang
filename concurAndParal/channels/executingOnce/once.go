package main

import (
	"fmt"
	"sync"
)

var on sync.Once

func setup()  {
	fmt.Println("Configuring everything")
}

func doStuff(wg *sync.WaitGroup){
	on.Do(setup)
	fmt.Println("Hello")
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++{
		wg.Add(1)
		go doStuff(&wg)
	}
	wg.Wait()
}
