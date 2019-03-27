package main

import "fmt"

func main() {
	a := F()
	a[0]()//3
	a[1]()//3
	a[2]()//3
}
func F() []func() {
	b := make([]func(), 3, 3)
	for i := 0; i < 3; i++ {
		b[i] = func() {
			fmt.Println(i)
		}
	}
	return b
}