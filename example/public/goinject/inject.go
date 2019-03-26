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
	Name         string
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
	if g.named == nil {
		g.named = make(map[string]*Object)
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
		if o.Name != "" {
			if g.named[o.Name] != nil {
				return fmt.Errorf("provided two instances named %s", o.Name)
			}
			g.named[o.Name] = o

		} else {
			if !o.embedded {
				if g.unnamedType[o.reflectType] {
					continue
				}
				g.unnamedType[o.reflectType] = true
			}
		}

		g.unnamed = append(g.unnamed, o)
	}
	return nil
}
func (g *Graph) Ensure() error {
	for _, o := range g.named {
		if err := g.ensureSub(o); err != nil {
			return err
		}
	}
	i := 0
	for {
		if i == len(g.unnamed) {
			break
		}
		o := g.unnamed[i]
		i++

		if err := g.ensureSub(o); err != nil {
			return err
		}
	}

	return nil
}
func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func (g *Graph) ensureSub(o *Object) error {
	fmt.Println("Populate")
	if !isStructPtr(o.reflectType) {
		return nil
	}
	for i := 0; i < o.reflectValue.Elem().NumField(); i++ {
		field := o.reflectValue.Elem().Field(i)
		fieldType := field.Type()
		fieldTag := o.reflectType.Elem().Field(i).Tag
		tag := parseTag(fieldTag)
		if tag == nil {
			continue
		}
		if !isNilOrZero(field, fieldType) {
			fmt.Println("zero")
			continue
		}
		if tag.Name != "" {
			existing := g.named[tag.Name]
			if existing == nil {
				return fmt.Errorf(
					"did not find object named %s required by field %s in type %s",
					tag.Name,
					o.reflectType.Elem().Field(i).Name,
					o.reflectType,
				)
			}

			if !existing.reflectType.AssignableTo(fieldType) {
				return fmt.Errorf(
					"object named %s of type %s is not assignable to field %s (%s) in type %s",
					tag.Name,
					fieldType,
					o.reflectType.Elem().Field(i).Name,
					existing.reflectType,
					o.reflectType,
				)
			}
			field.Set(reflect.ValueOf(existing.Value))
			continue
		}
		fmt.Println(fieldType.Kind(), o.reflectType.Elem().Field(i).Name, "bbbbbb")
		if fieldType.Kind() == reflect.Struct {

			if !tag.Inline {
				return fmt.Errorf(
					"inline struct on field %s in type %s requires an explicit \"inline\" tag",
					o.reflectType.Elem().Field(i).Name,
					o.reflectType,
				)
			}
			err := g.Provider(&Object{
				Value:    field.Addr().Interface(),
				embedded: o.reflectType.Elem().Field(i).Anonymous,
			})
			if err != nil {
				return err
			}
			continue
		}
		if fieldType.Kind() == reflect.Interface || fieldType.Kind() == reflect.Map {
			continue
		}
		newValue := reflect.New(fieldType.Elem())
		newObject := &Object{
			Value: newValue.Interface(),
		}

		// Add the newly ceated object to the known set of objects.
		err := g.Provider(newObject)
		if err != nil {
			return err
		}
		// Finally assign the newly created object to our field.
		field.Set(newValue)

	}
	return nil
}

func isNilOrZero(v reflect.Value, t reflect.Type) bool {
	switch v.Kind() {
	default:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(t).Interface())
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
}

type tag struct {
	Name   string
	Inline bool
}

func parseTag(t reflect.StructTag) *tag {
	if t == "" {
		return nil
	}
	r := regexp.MustCompile("^" + injectTag + ".*")
	if r.MatchString(string(t)) {
		name := t.Get(injectTag)
		switch name {
		case "":
			return &tag{}
		case "inline":
			return &tag{Inline: true}
		default:
			return &tag{Name: name}
		}
	}
	return nil
}
