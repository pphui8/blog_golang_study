# notes of learning golang

### [youtube video](https://www.youtube.com/watch?v=YS4e4q9oBaU)

---
# index
## README.md
> #### Variables
> #### Arrays and Slices 
> #### Maps and Structs
> #### If and Switch Statements 
> #### Looping 

---

## README2.md
> #### Defer, Panic, and Recover 
> #### Pointers 
> #### Functions 

---

## README3.md
> #### Interfaces 
> #### Goroutines 

---

## README4.md
> #### Channels

---

## variable (simillar to rust)
> 1. shadow in go:
> 2. you can`t declear same variable twice but you can shadow that in a smaller scope (ex: global to block)

> 3. every uninitialized variable was assgined by 0 (false)

> 4. rune is a alise to int32, to index string

> 5. visibility  
> 5.1 inside a block: private
> 5.2 outside a block: protect
> 5.3 outside a block with capitalized first letter: public

```go
// var name type = value
var i int = 10
var i, j int = 10, 20

// ## auto assert
// var name = value
var isTrue = true

// ## := equals to declear and assgin,
// ## = equals to only assgin
// name := value
// this syntax can only be used to declear a local variable
a := 10

// ## declear in a block (global variable mostly)
var (
    IP string = "xx.xx.xx.xx"
    hostname string = "xxx.com"
)
```
### convert variable (explicity only)
```go
var i float32 = 42.5
var j int32 = int32(i)
```

### iota
> A kind of auto iterator in enumrate
```go
const (
    SwitchOff = 0
    SwitchOn  = 1
)
// or simply
const (
    SwitchOff = iota
    SwitchOn
)
const (
    _ = iota    // ignore 0
    KB = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```
---

### array & slice
> size of array can`t be changed during the runtime
> so we can use slice instead of array
> or use make function to create a bigger array
```go
grades := [3]int{97, 85, 93}
grades := [...]int{97, 85, 93}

var students [3]string

// slice
a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
b := a[:]
c := a[3 : ]    // from 4th to the end
d := a[ : 6]    // from 0th to 6th
e := a[3 : 6]   // from 4th to 6th

// make function
// make(type, length, capacity)
// length indecate the length of value 0
a := make([]int, 3) // [0, 0, 0]
a := make([]int, 3, 100)

// append function
// result = append(target, value...)
a = append(a, 10)

// spread operation
a = append(a. []int{1, 2, 3}...)
```

---

## maps and structs
#### map
> 1. maps: key-value structure
> 2. order of map won`t be the same
> 3. passing map is passing pointer
```go
statePopulations := map[string]int{
    "California": 39250017,
    "Texas":      27862596,
    "Florida":    20612439,
    "New York":   19745289,
    "Pennsylvania":  15266476,
    "Ohio":       13607675,
}
// access
statePopulations["Ohio"]
// operations
statePopulations["Georgia"] = 10310371
delete(statePopulations, "Georgia")

// access sytax candy\
// ok stands for wether "California" is exsisting
pop, ok := statePopulations["California"]
// returns wether "California" is exsisting
_, ok := statePopulations["California"]
```

#### struct

```go
// declear
type Doctor struct {
	number int
	actorName string
	companions []string
}
// useage (protect)
aDoctor := Doctor{
    number: 1,
    actorName: "John Doe",
    companions: []string{"Jane Doe", "Jack Doe"},
}
// useage (public)
aDoctor := Doctor{
    Number: 1,
    ActorName: "John Doe",
    Companions: []string{"Jane Doe", "Jack Doe"},
}
// this sytax must match ther corresponding order
aDoctor := Doctor{
    1,
    "John Doe",
    []string{"Jane Doe", "Jack Doe"},
}
// sytax candy (anony)
aDoctor := struct{name string}{name: "Dr. Strange"}
```

#### inherit (by using composition via embedding)

``` go
type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	Animal  // or Animal Animal
	SpeedKPH float64
	CanFly bool
}
b := Bird{}
b.Name = "Emu"
b.Origin = "Australia"
b.SpeedKPH = 48
b.CanFly = false
```

#### tag (describe the value)

```go
type Animal struct {
	Name string `required max:"100"`
	Origin string
}
// get the tag message
t := reflect.TypeOf(Animal{})
field, _ := t.FieldByName("Name")
fmt.Println(field.Tag)
```

---

## if - else statement
> 1. initialize sytax
```go
if pop, ok := statePopulations["Florida"]; ok {
    fmt.Println(pop)
}
```

> 2. switch
```go
i := 10
switch {
    case i < 10:
        fmt.Println("less than 10")
        // if we wanna go ahead
        fallthrough
    case i < 20:
        fmt.Println("less than 20")
        // if we wanna break this ealry
        break
    default:
        fmt.Println("less than 10")
}
```

> 3. type switch
```go
switch i.(type) {
    case int:
        fmt.Println("i is a int")
    case float64:
        fmt.Println("i is a float64")
    case string:
        fmt.Println("i is a string")
    default:
        fmt.Println("i has another type")
}
```

---

## loop
> 1.1++ is not a pression, it`s a statement

> 2.```for i < 5 {``` stands for while loop
> or just ```for {``` stands for infinit loop

> 3.break with laber
```go
// define a laber here
Loop:
    for i := 1; i <= 3; i++ {
        for j := 1; i <= 3; i++ {
            fmt.Println(i * j)
            if i * j >= 3 {
                // break here
                break Loop
            }
        }
    }
```

> 4.loops for collection
```go
s := [3]int{1, 2, 3}
for k, v := range s {
    fmt.Println(k, v)
}
// or
s := "hello world"
for k, v := range s {
    fmt.Println(k, string(v))
}
```