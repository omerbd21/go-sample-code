package main

import "fmt"

var x int
var y string
var z bool

func main() {
	s := fmt.Sprintf("%d %v %t", x, y, z)
	fmt.Println(s)
}
