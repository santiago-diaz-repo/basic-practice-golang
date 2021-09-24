package main

import "fmt"

type Greeting struct {
	mockeableInterface MockeableInterface
}

func NewGreeting(mockeableInterface MockeableInterface) Greeting{
	return Greeting{
		mockeableInterface: mockeableInterface,
	}
}

func (g *Greeting) ExecuteGreeting(name string){
	fmt.Println(g.mockeableInterface.GetGreeting(name))
}

func main() {
	mockeable := mockeable{}
	greet := NewGreeting(&mockeable)
	greet.ExecuteGreeting("Test")
}

