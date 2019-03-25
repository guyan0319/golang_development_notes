package goinject

import (
	"fmt"
	"reflect"
	"sync"

	"regexp"
)

//    `inject:""`
//    `inject:"inline"`
//    `inject:"foo"`
var injectTag = "inject"
type Object struct {
	Value        interface{}
	name         string
	reflectType  reflect.Type
	reflectValue reflect.Value
	embedded     bool
}

type Graph struct {
	sync.Mutex
	unnamed     []*Object
	unnamedType map[reflect.Type]bool
	named       map[string]*Object
}

func (g *Graph) Provider(objects ...*Object) error {
	g.Lock()
	defer g.Unlock()
	if g.unnamedType == nil {
		g.unnamedType = make(map[reflect.Type]bool)
	}
	for _, o := range objects {
		o.reflectType = reflect.TypeOf(o.Value)
		o.reflectValue = reflect.ValueOf(o.Value)
		if !isStructPtr(o.reflectType) {
			return fmt.Errorf(
				"expected unnamed object value to be a pointer to a struct but got type %s "+
					"with value %v",
				o.reflectType,
				o.Value,
			)
		}
		if !o.embedded {
			if g.unnamedType[o.reflectType] {
				continue
			}
			g.unnamedType[o.reflectType] = true
		}
		g.unnamed = append(g.unnamed, o)
	}
	return nil
}
func (g *Graph) Ensure() error {
	i := 0
	for {

		if i == len(g.unnamed) {
			break
		}

		o := g.unnamed[i]
		i++

		if err := g.ensuresub(o); err != nil {
			return err
		}

	}
	return nil
}
func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func (g *Graph) ensuresub(o *Object) error {
	if !isStructPtr(o.reflectType) {
		return nil
	}
	//Loop:
	for i := 0; i < o.reflectValue.Elem().NumField(); i++ {
		field := o.reflectValue.Elem().Field(i)
		fieldType := field.Type()
		fieldTag := o.reflectType.Elem().Field(i).Tag
		tag := matchTag(fieldTag)
		fmt.Println(fieldTag.Get("inject"))
		//fieldTag := o.reflectType.Elem().Field(i).Tag.Get("inject")

		fmt.Println(fieldType,fieldTag,tag)
		//fieldName := o.reflectType.Elem().Field(i).Name
		//tag, err := parseTag(string(fieldTag))
		//if err != nil {
		//	return fmt.Errorf(
		//		"unexpected tag format `%s` for field %s in type %s",
		//		string(fieldTag),
		//		o.reflectType.Elem().Field(i).Name,
		//		o.reflectType,
		//	)
		//}

	}
	return nil
}
type tag struct {
	Name    string
	Inline  bool
	Sington bool
}
func matchTag(t reflect.StructTag) (*tag){
	if t =="" {
		return nil
	}
	r := regexp.MustCompile("^"+injectTag+".*")
	if r.MatchString(string(t)) {
		name:= t.Get(injectTag)
		switch name {
		case "":
			return &tag{Sington:true}
		case "inline":
			return &tag{Inline:true}
		default:
			return &tag{Name:name}
		}
	}
    return nil
}
//func parseTag(t string) (*tag, error) {
//found, value, err := structtag.Extract("inject", t)
//if err != nil {
//return nil, err
//}
//if !found {
//return nil, nil
//}
//if value == "" {
//return injectOnly, nil
//}
//if value == "inline" {
//return injectInline, nil
//}
//if value == "private" {
//return injectPrivate, nil
//}
//return &tag{Name: value}, nil
//}
