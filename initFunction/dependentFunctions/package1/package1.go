package package1

import "fmt"

func init() {
	fmt.Println("Executing init in package 1")
}

func DoStuff1()  {
	fmt.Println("Executing DoStuff in package 1")
}
