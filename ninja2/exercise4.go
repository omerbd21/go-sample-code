package main

import "fmt"

func main() {
	x := 100
	fmt.Printf("%d %b %x", x, x, x)
	y := 1 << x
	fmt.Printf("%d %b %x", y, y, y)
}
