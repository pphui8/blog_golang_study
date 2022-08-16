## interfaces
> 1. declear
```go
func main() {
	// implement the Writer interface here
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello, world.\n"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Print(string(data))
	return n, err
}
```

> 2. we could embed interfaces into interfaces
```go
func main() {
	// implement the Writer interface here
	var w ConsoleWriter = BufferWriter{}
	w.Write([]byte("Hello, world.\n"))
	w.Close()
}

type Writer interface {
	Write([]byte) (int, error)
}

type WriterCloseer interface {
	Close() error
}

type ConsoleWriter interface{
	Writer
	WriterCloseer
}

type BufferWriter struct{}

func (cw BufferWriter) Write(data []byte) (int, error) {
	n, err := fmt.Print(string(data))
	return n, err
}

func (cw BufferWriter) Close() error {
	return nil
}
```

> 3. interface conversion
> inplement different interface to do different behavier
```go
func main() {
    // under the above example
    // w has been converted into another type
    // it`s easy to get err here so..
    bwc, ok := w.(*BufferedWriterCloser)
    if ok {
        fmt.Prinln(bwc)
    } else {
        fmt.Prinln("conversion failed")
    }
    
}
```

> 4. different bewtten pointer receiver and concrate type receiver
> the concrate receiver can only use these functions that defined receive concrate type  
> but pointer receiver could use both (auto deref)

```go
func main() {
	var w Writer = &ConsoleWriter{}
	w.Write([]byte("Hello, world.\n"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

// pointer type rather then a concrate type
func (cw *ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Print(string(data))
	return n, err
}
```

#### tips
1. use many, small interfaces  
Single method interfaces are some of the most powerful and flexible
3. Don`t export interfaces for types that will be consumed
4. Do export interfaces for types that will used by package
5. Design functions and methods to receive interfaces whenever possible

---

## goroutines
gotoutine is a lightweight thread of execution rather than a os privided multi-

#### 1.create
```go
func main() {
	// create goroutine
	go hello()
	// wait for 1 second to see result
	time.Sleep(100 * time.Millisecond)
}

func hello() {
	fmt.Println("Hello World")
}
```

> 2. create by annonymous function
```go
func main() {
	var msg = "Hello World"
	// do not use variable outside this goroutine
	// pass this into the goroutine
	// do this ```go func(msg string) {```
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)
}

func hello() {
	fmt.Println("Hello World")
}
```

#### 2. wait group
> way to synchronize threads
> create a point for waiting the other thread compliate their job
```go
var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello World"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
}
```

#### 4. mutex
> another way to synchronize threads
> it provide a locked variable that makes sure only one thread is reacting with this variable

```go
var m = sync.Mutex{}
// or a read & write lock
var rwm = sync.RWMutex{}

func doSomething() {
	m.Lock()
	// or lock read only
	rwm.RLock()

	go aGoRoutine()

	m.RUnlock()
	rwm.RUnlock()
}
```

#### runtime.GOMAXPROCS()
set / get avaliable processes
> when get, pass -1 as argument
> runtime.GOMAXPROCS(-1)

#### prevent data race
compile with ```go run =race```