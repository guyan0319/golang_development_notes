package public

import (
	"fmt"
	"reflect"
	"strings"
)

func Explicit(v reflect.Value, depth int) {
	if v.CanInterface() {
		t := v.Type()
		switch v.Kind() {
		case reflect.Ptr:
			Explicit(v.Elem(), depth)
		case reflect.Struct:
			fmt.Printf(strings.Repeat("\t", depth)+"%v %v {\n", t.Name(), t.Kind())
			for i := 0; i < v.NumField(); i++ {
				f := v.Field(i)
				if f.Kind() == reflect.Struct || f.Kind() == reflect.Ptr {
					fmt.Printf(strings.Repeat("\t", depth+1)+"%s %s : \n", t.Field(i).Name, f.Type())
					Explicit(f, depth+2)
				} else {
					if t.Field(i).Name == "buf" {
						//fmt.Println(unsafe.Alignof(f))
						//p := (uintptr)(f.Pointer())
						//fmt.Println(f.UnsafeAddr())
						//fmt.Printf("%+p", f.UnsafeAddr())
						//p := (*int)(unsafe.Pointer(&f))
						//fmt.Println(f.Kind())

					}
					if f.CanInterface() {
						fmt.Printf(strings.Repeat("\t", depth+1)+"%s %s : %v \n", t.Field(i).Name, f.Type(), f.Interface())
					} else {
						fmt.Printf(strings.Repeat("\t", depth+1)+"%s %s : %v \n", t.Field(i).Name, f.Type(), f)
					}
				}
			}
			fmt.Println(strings.Repeat("\t", depth) + "}")
		}
	} else {
		fmt.Printf(strings.Repeat("\t", depth)+"%+v\n", v)
	}
}
