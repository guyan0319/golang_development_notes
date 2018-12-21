package main

import "fmt"

func main() {
	var m map[string]int
	m = map[string]int{"one": 1, "two": 2}
	n := m
	fmt.Printf("%p\n", &m) //0xc000074018
	fmt.Printf("%p\n", &n) //0xc000074020
	fmt.Println(m)         // map[two:2 one:1]
	fmt.Println(n)         //map[one:1 two:2]
	changeMap(&m)
	fmt.Printf("%p\n", &m) //0xc000074018
	fmt.Printf("%p\n", &n) //0xc000074020
	fmt.Println(m)         //map[one:1 two:2 three:3]
	fmt.Println(n)         //map[two:2 three:3 one:1]
}
func changeMap(m *map[string]int) {
	//m["three"] = 3 //这种方式会报错 invalid operation: m["three"] (type *map[string]int does not support indexing)
	(*m)["three"] = 3                    //正确
	fmt.Printf("changeMap func %p\n", m) //changeMap func 0x0
}
