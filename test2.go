package main

import (
	"fmt"
)

func goRoutineA(a <-chan string) {

	val := <-a
	fmt.Println("goRoutineA received the data", val)
}
func goRoutineB(a chan string, data string) {
	a <- data
	fmt.Println("goRoutineB send the data", data)
}

func main() {

	//ch := make(chan string)
	//go goRoutineB(ch, "hello")
	//go goRoutineA(ch)
	//time.Sleep(time.Second * 1)
}
