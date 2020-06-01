package main

import "fmt"
import "reflect"

func main(){
	var a interface{} = "asdf"

	// tn := a.(int)
	// fmt.Println(tn)

	fmt.Println(reflect.TypeOf(a))
}