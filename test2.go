package main

import (
	"fmt"
	"time"
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
	ch := make(chan string)

	fmt.Printf("%T \n", ch)
	fmt.Println(ch)
	//extractChan(reflect.ValueOf(ch))
	go goRoutineB(ch, "hello")
	go goRoutineA(ch)
	time.Sleep(time.Second * 1)
}

//func extractChan(v reflect.Value) (interface{}, error) {
//	if v.Kind() != reflect.Chan {
//		return nil, errors.New("invalid input")
//	}
//	var ch interface{}
//	ch = v.Interface()
//	return ch, nil
//}
