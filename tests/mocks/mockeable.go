package main

type mockeable struct {

}

func (m *mockeable) GetGreeting(name string) string{
	return string("This is a greet from a real object: " + name)
}