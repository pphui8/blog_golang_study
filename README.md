## variable (simillar to rust)
> 1. shadow in go:
> 2. you can`t declear same variable twice but you can shadow that in a smaller scope (ex: global to block)

> 3. every uninitialized variable was assgined by 0 (false)

> 4. rune is a alise to int32, to index string

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
### array
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

