package main

import "time"

func Prime(number int) bool {
	time.Sleep(300 * time.Millisecond)
	switch number {
	case 0:
		return false
	case 1:
		return false
	}
	var i int = 2
	for i < number {
		if number%i == 0 {
			return false
		}
		i++
	}
	return true
}


func main() {

}