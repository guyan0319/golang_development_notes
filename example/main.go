package main

import "fmt"

func main() {
	fmt.Printf("hello world!")
	//var s []string
	//s := []string{"dd", "fsdf", "f"}
	//	////s[3] = "dsf"
	//	//slice1 := s[1:3]
	//	////slice1 := slice1[:1]
	//	//slice1 = append(slice1, "aa", "bb")
	//	//fmt.Println(slice1)
	//s = append(s, "fsd")

	//s := []string{"123", "123"} //切片
	//var slice1 []string = make([]string, 2)
	//slice1 = append(slice1, "a", "fd", "fd", "fd", "fd", "fd", "fd", "fd", "fd", "fd", "fd", "fd", "fd")
	//fmt.Println(slice1)
	p := new([2]int)
	p[0] = 22
	b := make([]int, 10, 50)
	fmt.Println(p, b)
}
