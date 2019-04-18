package main

import "fmt"

func main() {
	var mv *map[string]string
	fmt.Printf("mv: %p %#v \n", &mv, mv)//mv: 0xc042004028 (*map[string]string)(nil)
	mv = new(map[string]string)
	fmt.Printf("mv: %p %#v \n", &mv, mv)//mv: 0xc000006028 &map[string]string(nil)
	(*mv) = make(map[string]string)
	(*mv)["a"] = "a"
	fmt.Printf("mv: %p %#v \n", &mv, mv)//mv: 0xc042004028 &map[string]string{"a":"a"}

}
