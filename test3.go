package main

import (
	"example/example/public"
	"fmt"
	"reflect"
)

type Users struct {
	Id   int
	Name string
	Age  int
	Ext  Info
}
type Info struct {
	Detail string
}

func (u Users) Login() {
	fmt.Println("login")
}

func main() {
	//ch := make(chan string, 3)
	//fmt.Println(ch)
	//vv := runtime.Getchan()
	u := &Users{Id: 12}
	//t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	//fmt.Printf("%+v\n", v)

	//t := reflect.TypeOf(ch)
	public.Examiner(v, 0)
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
