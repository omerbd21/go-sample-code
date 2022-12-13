package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() string {
	return "Hello, my name is " + p.first + " and i'm " + fmt.Sprint(p.age)
}
func main() {
	p := person{
		first: "Omer",
		last:  "Ben David",
		age:   22,
	}
	fmt.Println(p.speak())
}
