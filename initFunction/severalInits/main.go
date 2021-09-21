// The execution order is in the order of declaration
package main

import "fmt"

func init() {
	fmt.Println("Executing first init block")
}

func init() {
	fmt.Println("Executing second init block")
}

func init() {
	fmt.Println("Executing third init block")
}

func main() {
	fmt.Println("Main starts")
}