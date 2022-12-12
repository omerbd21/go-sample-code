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
	people := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}
	for _, v := range people {
		fmt.Println(v.first, v.last)
		for _, t := range v.favIceCream {
			fmt.Println(t)
		}
	}
}
