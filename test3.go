package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	ch <- 3
	i, isClose := <-ch
	if !isClose {

	}
	fmt.Println(<-ch)
}
