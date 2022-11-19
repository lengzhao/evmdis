package evmdis

import (
	"reflect"
)

type TypeMap struct {
	data map[reflect.Type]interface{}
}

func (tym *TypeMap) Get(obj interface{}) {
	element := reflect.ValueOf(obj).Elem()
	if value, ok := tym.data[element.Type()]; ok {
		element.Set(reflect.ValueOf(value))
	} else {
		element.Set(reflect.Zero(element.Type()))
	}
}

func (tym *TypeMap) Pop(obj interface{}) {
	element := reflect.ValueOf(obj).Elem()
	if value, ok := tym.data[element.Type()]; ok {
		element.Set(reflect.ValueOf(value))
		delete(tym.data, element.Type())
	} else {
		element.Set(reflect.Zero(element.Type()))
	}
}

func (tym *TypeMap) Set(obj interface{}) {
	element := reflect.ValueOf(obj).Elem()
	tym.data[element.Type()] = element.Interface()
}

func NewTypeMap() *TypeMap {
	return &TypeMap{make(map[reflect.Type]interface{})}
}
