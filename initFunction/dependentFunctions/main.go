// The execution order is:
// 1. init in package1
// 2. dependent function
// 3. init in main
package main

import (
	_ "basic-practice-golang/initFunction/dependentFunctions/package1"
	"fmt"
)

var firstVariable = previous()

func previous() int {
	fmt.Println("Executing dependent function")
	return 89
}

func init(){
	total := firstVariable + 10
	fmt.Printf("Executing init main: %d\n \n",total)
}

func main()  {
	fmt.Println("main starts")
}
