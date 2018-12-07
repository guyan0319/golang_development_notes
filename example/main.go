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
type Manager struct {
	User  //反射会将匿名字段作为一个独立字段来处理
	Title string
}

func (u User) Login() {
	fmt.Println("login")
}

func main() {
	m := Manager{User: User{1, "Jack", 12}, Title: "123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))                   //#号会将reflect的struct的详情页打印出来，可以看出来这是一个匿名字段
	fmt.Printf("%#v \n", t.FieldByIndex([]int{0, 0})) //此时 我们就可以将User当中的ID取出来,这里面需要传进方法中的是一个int类型的slice，User相对于manager索引是0，id相对于User索引也是0
	fmt.Printf("%v \n", t.FieldByIndex([]int{0, 1}))
	v := reflect.ValueOf(m)
	fmt.Printf("%#v\n", v.Field(0))
}
