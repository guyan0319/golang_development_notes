package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id     int
	Name   string
	Amount float64
}
type HandlerTypeVoid func()
type HandlerTypeString func() string
type HandlerTypeError func(interface{}) error

func main() {
	var i interface{}
	i = "string"
	fmt.Println(i)
	i = 1
	fmt.Println(i)
	i = User{Id: 2}
	//i.(User).Id = 15  //运行此处会报错，在函数中修改interface表示的结构体的成员变量的值，编译时遇到这个编译错误，cannot assign to i.(User).Id
	fmt.Println(i.(User).Id)
	i = test
	r := i.(func(v interface{}) error)("test_1")
	fmt.Println(r)
	//不同过反射调用函数
	var err error
	switch i.(type) { //通过使用.(type)方法可以利用switch来判断接口存储的类型。
	case func(string):
	case func(string, string):
		//...
	case func(interface{}) error:
		if f, ok := i.(func(v interface{}) error); ok {
			err = HandlerTypeError(f)("test_2")
		}
		break
	default:
		break
	}
	fmt.Println(err)
	//通过反射
	v := reflect.ValueOf(i)
	rargs := make([]reflect.Value, 1)
	rargs[0] = reflect.ValueOf("test_3")
	res := v.Call(rargs)
	fmt.Println(res)

}
func test(name interface{}) error {
	fmt.Println(name)
	return nil
}
