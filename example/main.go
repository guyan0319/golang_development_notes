package main

import (
	"fmt"
)

func main() {
	var a [1]int
	c := a[:]
	fmt.Println(c)
}
