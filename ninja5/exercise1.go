package main

import "fmt"

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
	p1 := person{
		first:       "Omer",
		last:        "Ben David",
		favIceCream: []string{"Peanut Butter"},
	}
	p2 := person{
		first:       "Roni",
		last:        "Ben David",
		favIceCream: []string{"Gum"},
	}
	fmt.Println(p1.first, p1.last)
	for _, v := range p1.favIceCream {
		fmt.Println(v)
	}
	fmt.Println(p2.first, p2.last)
	for _, v := range p2.favIceCream {
		fmt.Println(v)
	}
}
