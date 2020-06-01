package main

import "fmt"


var n int


func g(){
	fmt.Println(n)
}

func main(){
	x := []int{1,2,3,4}
	i := 0
	i, x[i] = 3, 10
	fmt.Println(i, x)


	fmt.Println(n)
	n, _ := 10, 2
	fmt.Println(n)
	g()

	a, b := 5, 6
	a, c := 7, 8
	fmt.Println(a, b, c)

	// a, c := 9 ,9 
	// fmt.Println(a, b, c)


}
