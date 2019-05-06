package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := map[string]string{"abc": "abc"}
	t := reflect.TypeOf(a)
	b := reflect.MakeMap(t)
	fmt.Println(b)

}
