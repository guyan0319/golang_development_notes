package main

import "fmt"

func main() {
	p := new([2]int)
	p[0] = 22
	b := make([]int, 10, 50) //第一个参数是类型，第二个参数是分配的空间，第三个参数是预留分配空间
	a := b[:cap(b)]
	fmt.Println(a)

}
