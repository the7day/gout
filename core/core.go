package core

import (
	"errors"
	"reflect"
	"unsafe"
)

type FormFile string

type FormMem []byte

var ErrUnknownType = errors.New("unknown type")

func LoopElem(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return v
		}
		v = v.Elem()
	}

	return v
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func NewPtrVal(defValue interface{}) interface{} {
	p := reflect.New(reflect.TypeOf(defValue))
	p.Elem().Set(reflect.ValueOf(defValue))
	return p.Interface()
}
