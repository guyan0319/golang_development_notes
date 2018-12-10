package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	//末尾插入值为1的元素，并返回该元素。
	v1 := l.PushBack(1)
	//首部插入值为2的元素，并返回该元素
	v2 := l.PushFront(2)
	//在元素v1前插入3
	l.InsertBefore(3, v2)
	//在元素v1后插入4
	l.InsertAfter(4, v1)

	fmt.Printf("len: %v\n", l.Len())
	fmt.Printf("first: %#v\n", l.Front())
	fmt.Printf("second: %#v\n", l.Front().Next())
	// 遍历列表并打印其内容。
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
