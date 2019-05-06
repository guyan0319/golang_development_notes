package public

import (
	"fmt"
	"reflect"
)

func Examiner(v reflect.Value, depth int) {
	t := v.Type()
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		Examiner(v.Elem(), 0)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			fmt.Printf("%s %s : %v \n", t.Field(i).Name, f.Type(), f.Interface())
		}
	}

}

// func Examiner(t reflect.Value, depth int) {
//	//fmt.Printf("%+v\n", t.Kind())
//	//fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
//	switch t.Kind() {
//	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
//		//fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
//		//fmt.Println(t)
//		//fmt.Println(t.Elem())
//		Examiner(t.Elem(), depth+1)
//	case reflect.Struct:
//		fmt.Printf("%v %v{\n", t.Name(), t.Kind())
//		for i := 0; i < t.NumField(); i++ {
//			f := t.Field(i)
//			v := reflect.ValueOf(f)
//			fmt.Println(v)
//			fmt.Println(strings.Repeat("\t", depth+1), f.Name, f.Type.Name())
//
//			//if f.Tag != "" {
//			//	fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
//			//	fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
//			//}
//		}
//		fmt.Println(strings.Repeat("\t", depth+1) + "}")
//	}
//}
