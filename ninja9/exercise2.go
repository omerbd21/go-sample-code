package main

import (
	"fmt"
)

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func (p *person) speak() {
	fmt.Println("I am " + p.First)
}
func main() {
	p1 := person{
		First: "Omer",
		Last:  "Ben David",
		Age:   22,
	}
	saySomething(&p1)
	//samSomething(p1) will not compile
}
