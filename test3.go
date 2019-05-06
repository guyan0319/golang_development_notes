package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func main() {
	ch := make(chan string)
	fmt.Println(ch)
	c := runtime.Getchan()
	t := reflect.TypeOf(c)
	examiner(t, 0)

	//chType := reflect.TypeOf(ch)

	//examiner(chType, 0)

	//f := test.Foo{}
	//fmt.Println(f)
	//fp := &f
	//
	//fType := reflect.TypeOf(f)
	//fpType := reflect.TypeOf(fp)
	//
	//examiner(fType, 0)
	//examiner(fpType, 0)

}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		fmt.Println(t.Elem())
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if isStructPtr(f.Type) {
				examiner(f.Type, 0)
			}

			//if f.Tag != "" {
			//	fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
			//	fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			//}
		}
	}
}
func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
