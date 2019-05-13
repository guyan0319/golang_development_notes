package main

import (
	"example/example/public"
	"fmt"
	"reflect"
	"runtime"
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
}
func main() {
	fmt.Println("1")
	//ch := make(chan int, 3)
	fmt.Println("2")
	ch := make(chan int)
	//fmt.Println(ch)
	//ch := make(chan int)
	//go goRoutineA(ch)
	//go goRoutineB(ch)
	//go goRoutineC(ch, 3)
	//go goRoutineC(ch, 4)
	//go goRoutineC(ch, 5)
	//go goRoutineC(ch, 6)
	//go goRoutineC(ch, 4)

	ch <- 3
	//ch <- 4
	//ch <- 5
	//ch <- 6
	//ch <- 7
	val := <-ch
	fmt.Println("goRoutineB  received the data", val)
	vv := runtime.Getchan()
	v := reflect.ValueOf(vv)
	public.Explicit(v, 0)
	//ww := runtime.Getwaitq()
	//w := reflect.ValueOf(ww)
	//public.Explicit(w, 0)
	//ch <- 3
	//ch <- 4
	//ch <- 5
	//ch <- 6
	//ch <- 7
	//ch <- 8
	//ch <- 9

	time.Sleep(time.Second * 1)
}
