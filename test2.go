package main

import (
	"fmt"
	"time"
)

func goRoutineA(a <-chan int) {
	val := <-a
	fmt.Println("goRoutineA received the data", val)
}
func goRoutineB(b chan int) {
	val := <-b
	fmt.Println("goRoutineB  received the data", val)
}
func goRoutineC(ch chan int, i int) {
	ch <- i
	ch <- 5
}
func main() {
	ch := make(chan int, 3)
	//ch := make(chan int)
	//go goRoutineA(ch)
	//go goRoutineB(ch)
	//go goRoutineC(ch, 3)
	//go goRoutineC(ch, 4)
	ch <- 3
	val := <-ch
	fmt.Println("goRoutineB  received the data", val)
	//ch <- 3
	//ch <- 4
	//ch <- 5
	//ch <- 6
	//ch <- 7
	//ch <- 8
	//ch <- 9

	time.Sleep(time.Second * 1)
}
