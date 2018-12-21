package main

import "fmt"

func main() {
	p := 5
	change(&p)
	fmt.Println("p=", p)
}
func change(p *int) {
	*p = 0
}
