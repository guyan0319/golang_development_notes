package goinject

import (
	"fmt"
	"reflect"
	"sync"
)

//    `inject:""`
//    `inject:"sington"`
//    `inject:"sington,demo"`
//    `inject:"demo"`

// An Object in the Graph.
type Object struct {
	Value        interface{}
	Name         string // Optional
	reflectType  reflect.Type
	reflectValue reflect.Value
	embedded     bool // If true, the Object is an embedded struct provided internally
	private      bool // If true, the Value will not be used and will only be populated

}

// The Graph of Objects.
type Graph struct {
	sync.Mutex
	unnamed     []*Object
	unnamedType map[reflect.Type]bool
	named       map[string]*Object
}

// Provide objects to the Graph. The Object documentation describes
// the impact of various fields.
func (g *Graph) Provide(objects ...*Object) error {
	for _, o := range objects {
		o.reflectType = reflect.TypeOf(o.Value)
		o.reflectValue = reflect.ValueOf(o.Value)

		if o.Name == "" {
			if !isStructPtr(o.reflectType) {
				return fmt.Errorf(
					"expected unnamed object value to be a pointer to a struct but got type %s "+
						"with value %v",
					o.reflectType,
					o.Value,
				)
			}
			if !o.private {
				if g.unnamedType == nil {
					g.unnamedType = make(map[reflect.Type]bool)
				}
				if g.unnamedType[o.reflectType] {
					return fmt.Errorf(
						"provided two unnamed instances of type *%s.%s",
						o.reflectType.Elem().PkgPath(), o.reflectType.Elem().Name(),
					)
				}
				g.unnamedType[o.reflectType] = true
			}
			g.unnamed = append(g.unnamed, o)

		} else {
			if g.named == nil {
				g.named = make(map[string]*Object)
			}
			if g.named[o.Name] != nil {
				return fmt.Errorf("provided two instances named %s", o.Name)
			}
			g.named[o.Name] = o
		}

	}
	return nil
}
func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
