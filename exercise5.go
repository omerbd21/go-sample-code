package main

import "fmt"

type marcobalbul int

var x marcobalbul
var y marcobalbul

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
}
