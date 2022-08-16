## channels
synchronize data between different multiple goroutine

> 1. channel could only storage one value a time
> 2. channel will copy the value rather than passing pointer

baisc ueage
```go
func main() {
	// create a channel
	ch := make(chan int)
	wg.Add(2)
	go func() {
		// revieving data from channel
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		i := 42
		// sending data to channel
		ch <- i
		fmt.Println(i)
		wg.Done()
	}()
    wg.Wait()
}
```

create a read only / write only channel
```go
func main() {
	ch := make(chan int)
	wg.Add(2)
	// declear a read only channel
	go func(ch <-chan int) {
		i := <- ch
		fmt.Println("abc", i)
		wg.Done()
	}(ch)
	// declear a write only channel
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}
```

### mutiple writer / sender
if we wanna sending one message but multiple receiver 
we will encount a dead lock like this
```go
go func(ch chan<- int) {
    ch <- 42
    // if thers`s only one receiver, process will be blocked here 
    // because none receive this 
    ch <- 27
}
```

##### 1. Buffered channel
declear a channel that receive mutiple messages  
not a idear method  
```go
ch := make(chan int, 50)
```

##### 2. Idealized method
to make a goroutine to not to block other goroutines
```go
func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		// deadlock here
		// keeping listening the channel but we stopped sending data
		for i := range ch {
			fmt.Println("Received:", i)
		}
		wg.Done()
	}(ch)
	go func() {
		i := 42
		ch <- i
		// close the channel to stop the receiving goroutine
		close(ch)
		wg.Done()
	}()
    wg.Wait()
}
```

#### look a little bit inside why receiver closed
```go
go func(ch <-chan int) {
		for {
            // if channel closed, ok isn`t true
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
	}(ch)
```

#### select
use receiver to shut the sender down

1. use defer func to close ch 
```go
defer func() {
    close(ch)
}()
```

2. common situation - select
> 1. change ch into a struct type without field (signal only channel)
> 2. add select statement
> 3. I can`t fully understand this bull shit

```go
var ch = make(chan struct{})

func main() {
	go func() {
		ch <- struct{}{}
	}()
	<-ch
}

func logger() {
	for {
        // sender will keep untill one of the receiver recevied
		select {
        // 
		case entry := <-logCh:
			fmt.Println(entry)
		case <-doneCh:
        // if message received
			return
		}
	}
}
```
