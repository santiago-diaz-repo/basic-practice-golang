package package2

import (
	_ "basic-practice-golang/initFunction/severalPackages/package1"
	"fmt"
)

func init()  {
	fmt.Println("Executing init in package 2")
}

func DoStuff2()  {
	fmt.Println("Executing DoStuff in package 2")
}