// genericity
// genericity
package main

import (
    "fmt"
    "reflect"
)

type GenericSlice struct {
    elemType   reflect.Type
    sliceValue reflect.Value
}

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (self *GenericSlice) Init(sample interface{}) {
    value := reflect.ValueOf(sample)
    self.sliceValue = reflect.MakeSlice(value.Type(), 0, 0)
    self.elemType = reflect.TypeOf(sample).Elem()
}

func (self *GenericSlice) Append(e interface{}) bool {
    if reflect.TypeOf(e) != self.elemType {
        return false
    }
    self.sliceValue = reflect.Append(self.sliceValue, reflect.ValueOf(e))
    return true
}

func (self *GenericSlice) ElemType() reflect.Type {
    return self.elemType
}

func (self *GenericSlice) Interface() interface{} {
    return self.sliceValue.Interface()
}

func TestReflect() {
    gs := GenericSlice{}
    gs.Init(make([]int, 0))
    fmt.Printf("Element Type: %s\n", gs.ElemType().Kind()) // => Element Type: int
    result := gs.Append(2)
    fmt.Printf("Result: %v\n", result)             // => Result: true
    fmt.Printf("sliceValue: %v\n", gs.Interface()) // => sliceValue: [2]
}

func (f *Foo) reflect() {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}

func TestReflect2() {
	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}

	f.reflect()
}
