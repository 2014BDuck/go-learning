package main

import (
	"fmt"
	"reflect"
)

// type MyType struct{
// 	Dada string
// }

// func main(){
// 	mt := MyType{}
// 	t := reflect.TypeOf(mt)
// 	fmt.Println(t.NumField())
// 	fmt.Println(t.Kind())
// 	fmt.Println(t.Method(0))
// }

// func (mt MyType) BFunc (x, y int) int{
// 	return 1
// }


type Student struct{
	Name string "学生姓名"
	Age int `a:"1111"b:"3333"`
}

func main(){
	s := Student{}
	rt := reflect.TypeOf(s)

	fieldName, ok := rt.FieldByName("Name")

	if ok{
		fmt.Println(fieldName.Tag)
	}

	fieldAge, ok2 := rt.FieldByName("Age")

	if ok2{
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}

	fmt.Println(rt.Name())
	fmt.Println(rt.NumField())
	fmt.Println(rt.PkgPath())
	fmt.Println(rt.String())

	fmt.Println("type.Kind().String()=",rt.Kind().String())
	fmt.Println("type.String()=",rt.String())
	for i := 0; i < rt.NumField(); i++{
		fmt.Println(rt.Field(i).Name)
	}




}