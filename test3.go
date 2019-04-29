package main

// Original Example from:
// https://www.socketloop.com/references/golang-reflect-select-and-selectcase-function-example
//
import (
	"fmt"
	"reflect"
)

func main() {

	// following replaces  " var sendCh := make(chan int) "
	type Foo struct {
		Ch chan int
	}
	sendCh := Foo{make(chan int)}
	// End of replacement code

	var increaseInt = func(c chan int) {
		for i := 0; i < 8; i++ {
			c <- i
		}
		close(c)
	}

	go increaseInt(sendCh.Ch)

	// This routine call replaces the code incorporated in "runJob"
	// It was done so I could call through an empty interface{}
	runJob(sendCh)
	// End replaement code--
}

func runJob(f interface{}) {
	var selectCase = make([]reflect.SelectCase, 1)

	// This code replaces just using the orginal "sendCh" value
	// I am trying here to construct "sendCh" from the interface value
	val := reflect.ValueOf(f).Elem()
	valueField := val.Field(0)
	sendCh := valueField.Interface()
	// End of replacement code

	selectCase[0].Dir = reflect.SelectRecv
	selectCase[0].Chan = reflect.ValueOf(sendCh)

	counter := 0
	for counter < 1 {
		chosen, recv, recvOk := reflect.Select(selectCase) // <--- here
		if recvOk {
			fmt.Println(chosen, recv.Int(), recvOk)

		} else {
			fmt.Println("Exit Condition Detected:  ", chosen, recv.Int(), recvOk)
			counter++
		}
	}
}
