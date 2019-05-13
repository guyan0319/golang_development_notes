package main

import (
	"fmt"
	"time"
)

func goRoutineD(ch chan int, i int) {
	time.Sleep(time.Second * 4)
	ch <- i
}
func goRoutineE(chs chan string, i string) {
	time.Sleep(time.Second * 3)
	chs <- i

}

func main() {
	ch := make(chan int, 5)
	chs := make(chan string, 5)

	go goRoutineD(ch, 5)
	go goRoutineE(chs, "ok")
	select {
	case msg := <-ch:
		fmt.Println(" received the data ", msg)
	case msgs := <-chs:
		fmt.Println(" received the data ", msgs)
		//default:
		//	fmt.Println("no data received ")
		//	time.Sleep(time.Second * 1)
	}

}
