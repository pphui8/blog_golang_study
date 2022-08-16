package main

import "fmt"

var ch = make(chan struct{})

func main() {
	go func() {
		ch <- struct{}{}
	}()
	<-ch
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Println(entry)
		case <-doneCh:
			return
		}
	}
}