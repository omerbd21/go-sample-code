package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) changeMe(first string, last string) {
	p.first = first
	p.last = last

}

func main() {
	p := person{
		first: "Omer",
		last:  "Ben David",
	}
	fmt.Println(p.first, p.last)
	p.changeMe("Sahar", "Shomroni")
	fmt.Println(p.first, p.last)

}
