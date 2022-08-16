package main

import "fmt"

func main() {
	g := greeter {
		greeting: "Hello",
		name: "World",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name string
}
 
// invoked by g.geet()
// g stands for the name of greeter
// so this is actually a method that given to the greeter struct
func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
}