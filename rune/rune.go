package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {

	a1 := '🦍'
	var a2 = 'k'
	var a3 rune = '🦡'

	a3+=1

	fmt.Printf("%c - %s\n", a1, reflect.TypeOf(a1))
	fmt.Printf("%c - %s\n", a2, reflect.TypeOf(a2))
	fmt.Printf("%c - %s\n", a3, reflect.TypeOf(a3))


	msg := "One 🐜"
	n1 := len(msg)
	n2 := utf8.RuneCountInString(msg)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println("-------")
	for _,v := range msg{
		fmt.Print(v)
		fmt.Println(" - "+string(v))
	}
}
