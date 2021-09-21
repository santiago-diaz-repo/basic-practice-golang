// In this scenario package2 has the package1 as dependency, however init in package1 is never executed
// The execution order is:
// 1. init in package2
// 2. init in main
package main

import (
	"basic-practice-golang/initFunction/severalPackages/package2"
	"fmt"
)

func init()  {
	fmt.Println("Executing init main")
}

func main()  {
	fmt.Println("main stars")
	package2.DoStuff2()
}