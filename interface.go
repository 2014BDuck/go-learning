package main

import (
	"fmt"
	// "io"
)

type Inter interface{
	Ping()
	Pang()
}

type Anter interface{
	Inter
	String()
}

type St struct{
	// Name string
}

func (St) Ping(){
	fmt.Println("Ping")
}

func (*St) Pang(){
	fmt.Println("Pang")
}

// func (St) String(){}

func main(){
	var st *St = nil
	var it Inter = st

	fmt.Printf("%p\n", st)
	fmt.Printf("%p\n", it)

	if it != nil{
		it.Pang()
		it.Ping()
	}
	// st := St{"Duck"}
	// fmt.Println(st)

	// var i interface{} = &st

	// if _, ok:= i.(Anter); ok{
	// 	fmt.Println("i fulfilled Inter")
	// }else{
	// 	fmt.Println("No!")
	// }

	// var j Inter
	// j = &st
	// switch v := j.(type){
	// case nil:
	// 	fmt.Println("asf")
	// default:
	// 	fmt.Println(v)
	// }
}