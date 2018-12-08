package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Login(name string) {
	fmt.Println("login" + name)
}
func (u User) LoginOut(name, name1 string) {
	fmt.Println("loginout" + name)
	fmt.Println("loginout" + name1)
}

func main() {
	user := User{Id: 1, Name: "jerry", Age: 29}
	val := reflect.ValueOf(&user) //获取Value类型，也可以使用reflect.ValueOf(&user).Elem()

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("herry")
	val.MethodByName("Login").Call(params) //通过名称调用方法
	paramstwo := make([]reflect.Value, 2)
	paramstwo[0] = reflect.ValueOf("herry")
	paramstwo[1] = reflect.ValueOf("jack")
	fmt.Println(params)
	val.Method(1).Call(paramstwo) //通过方法索引调用，paramstwo 含有两个参数

}
