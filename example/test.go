package main

import (
	"fmt"
	"time"
)

func main() {

	go fmt.Println("fasdf")
	c := 0
	start := time.Now()
	for i := 1; i < 2000000000; i++ {
		//c = (func() int {
		//	a := 1
		//	b := 3
		//	return a + b
		//})()
		c = F()
	}
	fmt.Println(c)
	t0 := time.Now()
	fmt.Printf("Cost time %v\n", t0.Sub(start))
}

func F() int {
	a := 1
	b := 3
	return a + b
}
