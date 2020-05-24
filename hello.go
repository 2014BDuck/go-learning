package main

import "fmt"

func main(){
	var arr [100]int
	fmt.Println(arr)

	arr2 := [...]float64{1:12.23, 10:2.0, 2:3}
	fmt.Println(arr2)
	fmt.Println(arr2[:5])


	a := make([]int, 1)
	a = append(a, 1,2,3,4,5)
	c := make([]int, 1)
	c = append(c, 8,9,8)
	fmt.Println(a, c)
	copy(a, c)
	fmt.Println(a, c)

	s := "heeloworls"
	sa := []rune(s)
	sb := []byte(s)
	fmt.Println(sa, sb)


	ma := map[string]*string{}
	ma2 := make(map[string]int)
	w := "world"
	w1 := "hehe"
	ma["hellp"] = &w
	ma["key2"] = &w1
	fmt.Println(ma)
	fmt.Println(ma)
	fmt.Println(ma)
	fmt.Println(ma)
	fmt.Println(ma)
	for k, v := range ma{
		fmt.Println(k, v)
	}

	for k, v := range ma{
		fmt.Println(k, v)
	}

	for k, v := range ma{
		fmt.Println(k, v)
	}
	fmt.Println(ma2)


	type Person struct{
		Name string
		age uint8
	}

	type Student struct{
		Person
		class string
	}

	a12 := Person{
		Name: "jiekun",
		age: 24,
	}

	a13 := Student{
		Person: a12,
		class: "LCS",
	}


	fmt.Println(a12)
	fmt.Println(a13)

	if x, y := sum(6,3); y< x{
		fmt.Println(y)
	}else{
		fmt.Println(x)
	}

	for i := range [7]int{1,2,3,4,5, 6, 7}{
		switch{
		case i > 3 && i < 5:
			fmt.Println("i bigger than 3!")
		case i <= 3:
			fmt.Println("i smaller than 3!")
		default:
			fmt.Println("Other")
		}
	}


}


func sum(a, b int) (int, int) {
	return a, b
}