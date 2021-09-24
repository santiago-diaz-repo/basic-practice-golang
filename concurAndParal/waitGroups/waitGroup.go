package main

import (
	"fmt"
	"sync"
)

var cache = make(map[int]string)


func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++{
		wg.Add(2)
		go searchInfo(i, &wg)
		go writeInfo(i, &wg)
	}
	wg.Wait()
	fmt.Println("It's time to finish")
}

func searchInfo(index int, wg *sync.WaitGroup)  {
	data := cache[index]
	fmt.Println("Data:" + data + "-")
	wg.Done()
}

func writeInfo(index int, wg *sync.WaitGroup){
	if cache[index] == ""{
		//cache[index] = strconv.Itoa(index)
	}
	wg.Done()
}
