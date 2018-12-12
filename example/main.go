package main

import "fmt"

func main() {
	//multiParam 可以接受可变数量的参数
	multiParam("jerry", 1)
	multiParam("php", 1, 2)
}
func multiParam(name string, args ...int) {
	fmt.Println(name)
	//接受的参数放在args数组中
	for _, e := range args {
		fmt.Println(e)
	}
}
