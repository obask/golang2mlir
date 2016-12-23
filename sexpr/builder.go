package sexpr

import (
	"reflect"
	"fmt"
	"go/ast"
	"go/token"
)

func _fake() {
	ast.Print(nil, nil)
	token.NewFileSet()
	reflect.ValueOf(nil)
	fmt.Print()
}


func CreateStruct(t reflect.Type, fields []reflect.Value) reflect.Value {
	//fmt.Println("CreateStruct: ", t)
	if t.Kind() != reflect.Struct {
		panic(t.Kind())
	}
	item := reflect.New(t).Elem()
	pos := 0
	n := t.NumField()
	for i := 0; i < n; i++ {
		// exclude non-exported fields because their
		// values cannot be accessed via reflection
		dst := item.Field(i)
		if (!IsBadField(dst)) {
			src := fields[pos]
			if !src.IsValid() {
				dst.Set(reflect.Zero(dst.Type()))
				pos++
			} else if src.Kind() == reflect.Ptr ||
					src.Kind() == reflect.Interface ||
					src.Type() == reflect.TypeOf(token.VAR) ||
					reflect.Bool == src.Kind() {
				//
				dst.Set(src)
				pos++
			} else if dst.Kind() == reflect.Slice {
				slice := reflect.MakeSlice(dst.Type(), 0, 0)
				length := src.Len()
				for j := 0; j < length; j++ {
					slice = reflect.Append(slice, src.Index(j))
				}
				//fmt.Println("DEBUG SOURCE:", src)
				//fmt.Println("DEBUG SOURCE:", slice)

				dst.Set(slice)
				pos++
			} else {
				panic(src)
			}
		}
	}
	return item.Addr()
}

func CreateSlice(fields []reflect.Value) reflect.Value {
	if len(fields) == 0 {
		return reflect.ValueOf(nil)
	}
	t := fields[0].Type()
	//fmt.Println("CreateSlice: ", t)

	if t.Kind() != reflect.Ptr && t.Kind() != reflect.Interface {
		panic(t)
	}
	slice := reflect.MakeSlice(reflect.SliceOf(t), 0, 0)
	for _,elem := range fields {
		slice = reflect.Append(slice, elem)
	}
	//fmt.Println("DEBUG SLICE:", slice)
	return slice
}
