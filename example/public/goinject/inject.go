package goinject

import "reflect"

//    `inject:""`
//    `inject:"sington"`
//    `inject:"sington,demo"`
//    `inject:"demo"`

// An Object in the Graph.
type Object struct {
	Value        interface{}
	Name         string             // Optional
	Fields       map[string]*Object // Populated with the field names that were injected and their corresponding *Object.
	reflectType  reflect.Type
	reflectValue reflect.Value
	created      bool // If true, the Object was created by us
	embedded     bool // If true, the Object is an embedded struct provided internally
}

// The Graph of Objects.
type Graph struct {
	unnamed []*Object
	named   map[string]*Object
}
