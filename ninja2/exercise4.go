package main

import "fmt"

const (
	x     = 5
	y int = 10
)

func main() {
	x := 100
	fmt.Printf("%d %b %x", x, x, x)
	y := 1 << x
	fmt.Printf("%d %b %x", y, y, y)
}
