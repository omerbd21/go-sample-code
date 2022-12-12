package main

import (
	"fmt"
)

func main() {
	x := 1913
	switch {
	case x == x:
		fmt.Println(x)
	case x != x:
		fmt.Println("No")
	}
}
