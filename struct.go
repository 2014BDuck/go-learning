package main

import (
	"fmt"
)

type Person struct{
	Name string
	Age int
}



// type Duck struct{
// 	string // Anonymous Field
// }


type T struct{
	a int
}

func (t T) Get() int{
	return t.a
}

func (t *T) Set(i int){
	t.a = i
}

func (t *T) Print(){
	fmt.Printf("%p, %v, %d \n", t, t, t.a)
}

type Int int

func (i *Int)Set(k Int){
	*i = k
}

func (i Int)Get()Int{
	return i
}

type Data struct{}


func (*Data) TestPointer(){}

func (Int)TestValue(){}

func (*Int)TestPointer(){}



type Animal struct{
	Name string
}

type Duck struct{
	Animal *Animal
	Speed int
}

type Pig struct{
	Animal
	Weight int
}

func (animal Animal)GetName(){
	fmt.Println(animal.Name)
}

func (pig Pig)Get(){
	fmt.Printf("%s weight is: %d\n", pig.Name, pig.Weight)
}

func main(){
	animal := Animal{"PeiQi"}
	p := Pig{animal, 888}

	fmt.Println(p)
	p.Get()
	p.GetName()
	// animal.Get()


	(*Data)(&struct{}{}).TestPointer()
	// (Data)(struct{}{}).TestPointer()

	var i Int = 10
	var j Int = 20
	i.Set(j)
	fmt.Println(i)


	fmt.Println((&i).Get())


	i.TestValue()

	(&i).TestPointer()
	(&i).TestValue()
	(Int(10)).TestValue()
	// (Int(10)).TestPointer()

	a := Person{"B.Duck", 5}
	b := Person{
		Name: "Rice",
		Age: 5,
	}

	c := new(Person)
	c.Name = "Egg"

	fmt.Println(a,b,*c)

	// d := new(Duck)
	// name = "Jack"
	// d.Name = &name
	// fmt.Prinln(d)
	fmt.Println(a.GetAge())
	fmt.Println(c.GetPointerAge())

	e := Duck{
		Animal: &Animal{"Jack"},
		Speed: 999,
	}
	fmt.Println(e)
	// fmt.Println(e.Animal.GetName())

	t := new(T)
	fmt.Println(t)


	f := t.Set
	f(2)
	t.Print()

	f(3)
	t.Print()

	fmt.Println(t.Get())
}

func (p Person)GetAge()int{
	return p.Age
}

func (p *Person)GetPointerAge()Person{
	return *p
}

// func (self Animal)GetName() string{
// 	return self.Name
// }