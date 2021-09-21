package package2

import "fmt"

func init()  {
	fmt.Println("Executing init in package 2")
}

func DoStuff2()  {
	fmt.Println("Executing DoStuff in package 2")
}