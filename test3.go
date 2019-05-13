package main

import (
	"fmt"
)

func goRoutineD(ch chan int, i int) {
	for   i := 1; i <= 5; i++{
		//time.Sleep(time.Second * 1)
		ch <- i
	}

}
func chanRange(chanName chan int) {
	for e := range chanName {
		fmt.Printf("Get element from chan: %d\n", e)
	}
}
func main() {
	ch := make(chan int, 5)
	go goRoutineD(ch, 5)
	chanRange(ch)

}
