package main


import (
	"fmt"
)
func main(){
	a1 := "hehe"
	a2 := "broken"
	a3 := "test"
	// test(&a1, &a2, &a3)
	do(test, &a1, &a2)
	fmt.Println(a1, a2, a3)
	do(test2, &a1, &a3)
	fmt.Println(a1, a2, a3)
	do2 := do
	do2(test, &a1, &a3)
	fmt.Println(a1, a2, a3)
	fmt.Printf("%T\n", test)

	
	fmt.Println(do3(1,3, plus))



	defer func(){
		fmt.Println("first")
	}()

	defer func(){
		panic("Made")
		fmt.Println("haha")
	}()

	fmt.Println("last?")

	// closure
	f := closure(1)
	fmt.Println(f(5))
	fmt.Println(f(5))

	g := closure(1)
	fmt.Println(g(5))
	fmt.Println(g(5))

	// panic

	defer catch()
	panic("Oh!")

}


func catch(){
	fmt.Println("recovering")
	recover()
}


func closure(a int) func(int) int{
	return func(i int) int{
		fmt.Println(a)
		a = a + i
		return a
	}
}


func plus(a, b int) int{
	return a+b
}
func test(arr ...*string){
	for _, each_var := range arr{
		*each_var = "fixed"
		// fmt.Println(each_var)
	}
	return
}

func test2(some_arr ...*string){
	return
}

type my_func func(...*string)

func do(f my_func, a *string, b *string){
	f(a, b)
	return
}


func do3(a, b int, f func(int, int)int)int{
	return f(a, b)
}