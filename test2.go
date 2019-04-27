package main

import "fmt"

func main() {
	//trace.Start(os.Stderr)
	//defer trace.Stop()
	// create new channel of type int
	ch := make(chan int, 1)
	fmt.Printf("%v\n", &ch)
	fmt.Printf("%+v\n", ch)
	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read from channel
	<-ch
}
