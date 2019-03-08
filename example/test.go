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

func (u User) Login() {
	fmt.Println("login")
}

func main() {
	u := User{1, "jerry", 23}
	t := reflect.TypeOf(u) //反射出一个interface{}的类型,main.User
	v := reflect.ValueOf(u)
	for i := 0; i < t.NumField(); i++ { //通过索引来取得它的所有字段，这里通过t.NumField来获取它多拥有的字段数量，同时来决定循环的次数
		f := t.Field(i)               //通过这个i作为它的索引，从0开始来取得它的字段
		val := v.Field(i).Interface() //通过interface方法来取出这个字段所对应的值
		fmt.Printf("%6s:%v =%v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ { //这里同样通过t.NumMethod来获取它拥有的方法的数量，来决定循环的次数
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}

	fmt.Println(t.Name())          //类型名 User
	fmt.Println(t.Kind().String()) //Type类型表示的具体分类  struct
	fmt.Println(t.PkgPath())       //反射对象所在的短包名 main
	fmt.Println(t.String())        //包名.类型名  main.User
	fmt.Println(t.Size())          //要保存一个该类型要多少个字节 32
	fmt.Println(t.Align())         //返回当从内存中申请一个该类型值时，会对齐的字节数 8
	fmt.Println(t.FieldAlign())    //返回当该类型作为结构体的字段时，会对齐的字节数 8

	fmt.Println(t.AssignableTo(reflect.TypeOf(u)))  // 如果该类型的值可以直接赋值给u代表的类型，返回真 true
	fmt.Println(t.ConvertibleTo(reflect.TypeOf(u))) // 如该类型的值可以转换为u代表的类型，返回真 true

	fmt.Println(t.NumField())             // 返回struct类型的字段数（匿名字段算作一个字段），如非结构体类型将panic  3
	fmt.Println(t.Field(0).Name)          // 返回struct类型的第i个字段的类型，如非结构体或者i不在[0, NumField())内将会panic  Id
	fmt.Println(t.FieldByName("Age"))     // 返回该类型名为name的字段（会查找匿名字段及其子字段），布尔值说明是否找到，如非结构体将panic
	fmt.Println(t.FieldByIndex([]int{0})) // 返回索引序列指定的嵌套字段的类型，等价于用索引中每个值链式调用本方法，如非结构体将会panic

	//ValueOf
	fmt.Println(v.IsValid()) //返回v是否持有值，如果v是value零值会返回假，此时v除了IsValid String Kind之外的方法都会导致panic
	fmt.Println(v.Kind())    //返回v持有值的分类，如果v是value零值，返回值为invalid  struct
	fmt.Println(v.Type())    //返回v持有值的类型Type表示  main.User

	//结构体指针
	//vv := &v
	fmt.Println(v.Convert(reflect.TypeOf(u)).FieldByName("Name")) // //转换为其他类型的值,如果无法使用标准Go转换规则来转换，那么panic  jerry
	pv := reflect.ValueOf(&u)
	pp := pv.Elem() //返回持有的接口的值，或者指针的值，如果不是interface{}或指针会panic,实际上是从 *User到User
	fmt.Println(pp)
	fmt.Println(v.FieldByName("Name").CanSet())  //是否可以设置Name的值  false
	fmt.Println(pp.FieldByName("Name").CanSet()) //是否可以设置Name的值  true
	pp.FieldByName("Name").SetString("newname")  //设置Name的值
	fmt.Println(u)                               //{1 newname 23}

	fmt.Println(pp.FieldByName("Name").Interface()) //把Name当做interface{}值  newname
	fmt.Println(pp.FieldByName("Name").String())    //返回v持有的值的字符串表示，如果v的值不是string也不会panic  newname
	var x int64
	fmt.Println(v.FieldByName("Age").OverflowInt(x)) //如果v持有值的类型不能溢出的表示x，会返回真，如果v的kind不是int int8-int64会panic false
}
