package main

import (
	"fmt"
	"strconv"
	"sync"
)

var cache = make(map[int]string)


func main() {
	var wg sync.WaitGroup
	m := &sync.RWMutex{}
	for i := 0; i < 10; i++{
		wg.Add(2)
		go searchInfo(i, &wg, m)
		go writeInfo(i, &wg, m)
	}
	wg.Wait()
	fmt.Println("It's time to finish")
}

func searchInfo(index int, wg *sync.WaitGroup, m *sync.RWMutex)  {
	m.RLock()
	data := cache[index]
	m.RUnlock()
	fmt.Println("Data:" + data + "-")
	wg.Done()
}

func writeInfo(index int, wg *sync.WaitGroup, m *sync.RWMutex){
	m.Lock()
	cache[index] = strconv.Itoa(index)
	m.Unlock()
	wg.Done()
}
