package main

import "fmt"


type Human struct{
	Name string
	Age int
}

type Student interface{
	GetName() string
	GetAge() int
}

func (s Human)GetName()string{
	return s.Name
}

func (s Human)GetAge()int{
	return s.Age
}

func main(){
	var i Student

	i = Human{"hehe", 8}

	fmt.Println(i.GetName())

	j := i.(Human)
	fmt.Println(j)
}