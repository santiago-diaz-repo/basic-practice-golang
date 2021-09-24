package main

import (
"fmt"
	"strconv"
	"sync"
)

var cache = make(map[int]string)


func main() {
	var wg sync.WaitGroup
	m := &sync.Mutex{}
	for i := 0; i < 10; i++{
		wg.Add(2)
		go searchInfo(i, &wg, m)
		go writeInfo(i, &wg, m)
	}
	wg.Wait()
	fmt.Println("It's time to finish")
}

func searchInfo(index int, wg *sync.WaitGroup, m *sync.Mutex)  {
	m.Lock()
	data := cache[index]
	m.Unlock()
	fmt.Println("Data:" + data + "-")
	wg.Done()
}

func writeInfo(index int, wg *sync.WaitGroup, m *sync.Mutex){
	m.Lock()
	if cache[index] == ""{
		cache[index] = strconv.Itoa(index)
		m.Unlock()
	}
	wg.Done()
}
