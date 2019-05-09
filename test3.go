package main

import (
	"example/example/public"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

type Users struct {
	Id     int
	Name   string
	Age    int
	Market map[int]string
	Source *Sfrom
	Ext    Info
}
type Info struct {
	Detail string
}
type Sfrom struct {
	Area string
}

func (u Users) Login() {
	fmt.Println("login")
}

func main() {

	ch := make(chan int, 3)
	fmt.Println(ch)
	vv := runtime.Getchan()
	v := reflect.ValueOf(vv)

	public.Explicit(v, 0)
	//fmt.Printf("%+v", v)
	//m := map[int]string{1: "abc"}
	//s := &Sfrom{Area: "beijing"}
	//i := Info{Detail: "detail"}
	//u := &Users{Id: 12, Market: m, Ext: i, Source: s}
	////t := reflect.TypeOf(u)
	//v := reflect.ValueOf(u)
	//public.Explicit(v, 0)
	//fmt.Printf("%+v\n", v)

	//t := reflect.TypeOf(ch)
	//public.Examiner(v, 0)
	//public.Examiner(t, 0)
	//t := reflect.TypeOf(u).Elem() //反射出一个interface{}的类型,main.User
	//v := reflect.ValueOf(u).Elem()
	//for i := 0; i < t.NumField(); i++ { //通过索引来取得它的所有字段，这里通过t.NumField来获取它多拥有的字段数量，同时来决定循环的次数
	//	f := t.Field(i)               //通过这个i作为它的索引，从0开始来取得它的字段
	//	val := v.Field(i).Interface() //通过interface方法来取出这个字段所对应的值
	//
	//	fmt.Printf("%v:%+v =%v\n", f.Name, f.Type, val)
	//}

}

func fun1() {
	//a := 2
	//c := (*string)(unsafe.Pointer(&a))
	//*c = "44"
	//fmt.Println(c)
	//fmt.Println(*c)
	a := "22"
	c := (*int)(unsafe.Pointer(&a))
	*c = 44
	fmt.Println(c)
	fmt.Println(&c)
	b := &c
	fmt.Println(&b)
}

func fun2() {
	a := "654"
	c := (*string)(unsafe.Pointer(&a))
	*c = "44"
	fmt.Println(c)
	fmt.Println(*c)
}
func fun3() {
	a := 3
	c := *(*string)(unsafe.Pointer(&a))
	c = "445"
	fmt.Println(c)

}
