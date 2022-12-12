package main

import "fmt"

type vehcile struct {
	doors int
	color string
}

type truck struct {
	v         vehcile
	fourWheel bool
}
type sedan struct {
	v      vehcile
	luxury bool
}

func main() {
	t := truck{
		v:         vehcile{doors: 2, color: "black"},
		fourWheel: true,
	}
	s := sedan{
		v:      vehcile{doors: 3, color: "white"},
		luxury: true,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.v.color, s.v.color)
}
