// The execution order depends on the declaration order, unless there are dependencies among packages:
// 1. init in package1
// 2. init in package2
// 2. init in main
package main

import (
	_ "basic-practice-golang/initFunction/severalPackages/package1"
	_ "basic-practice-golang/initFunction/severalPackages/package2"
	"fmt"
)

func init()  {
	fmt.Println("Executing init main")
}

func main()  {
	fmt.Println("main starts")
}