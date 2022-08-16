## Defer
> 1.declear a function that execute just before the function return

> 2.defer storage functions likes a stack
```go
// print:
// end
// middle
// start
func main {
    defer fmt.Println("start")
    defer fmt.Println("middle")
    defer fmt.Println("end")
}
```
> when to use?
> commonly to close http connet to pervent fotget to close the connetion
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
    // close after we got return value
    defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
```

---

## Panic && Recover
> 1.Panic: same to rust
```go
func panicFun() {
    panic("some thing went wrong")
    return
}
```

> 2.defer function still will be called after a panic occured

> 3.`Recover`: 
> 3.1. Recover could only be functional inside a defer function
> 3.2. if the goroutine paniced, recover will catch the panic value and execute again

> 4.`Recover` will return nil if application isn`t paniking
``` go
func main() {
	fmt.Println("start")
    // main function still gose well
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
    // deal with the panic down below
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
            // if we can`t handle the panic
            panic(err)
		}
	}()
	panic("I am panicking")
}
```

---

## Pointers
> basic ueage
```go
func main() {
	var ms *myStruct
	ms = new(myStruct)
	(*ms).foo = 42
    // compiler will deref this automaticlly
    // ms.foo = 42  // works as well
	fmt.Println((*ms).foo)
}

type myStruct struct {
	foo int
}
```

> 1.copy array is a copy
> so change b won`t change a
```go
a := [3]int{1, 2, 3}    // a is a array
b := a
```
> 2.slice doesn`t contain the data itself
> change b will change a
```go
a := []int{1, 2, 3}    // a is a array
b := a
```

---

## Function
> 1.upper case/lower case descide visibility

> 2.camml naming

> 3.sytax sugar
```go
func a(a string, b string) {
// or simply: func a(a, b string)

}
```

> 4.recive mutiple value
```go
func sum(values ...int) {

}
```

> 5.return pointer would be fine because golang won`t clean that

> 6.define return value name on declear
```go
func a() (result int) {
	result = 10
	// return the result
	return
}
```

> 7.return err
```go
func devide(a, b float64) (int, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorg("Cannot divide by zero")
	}
	return a / b, nil
}
```

> 8.annonyous funcion
```go
func main() {
	func() {
		// do something
	}()
	// or passing some value
	func(i int) {
		// do something
	}(i)
}
```

> 9.variable function
```go
f := func() {
	// do something
}

// or like js
var f func() = func() {
	// do something
}

// or spreate declear
var devide func(float64, float64) (float64, error)
devide = func(a, b float64) (float64, error) {
	// do something
}
```

> 10.make a function method
```go
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
// g stands for the name of greeter to provde a context
// g is a copy not a pointer
// so this is actually a method that given to the greeter struct
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
// or make this a pointer
func (g *greeter) greet() {
	// auto dereference
	fmt.Println(g.greeting, g.name)
}
```