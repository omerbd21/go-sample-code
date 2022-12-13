package main

import (
	"fmt"
	"math"
)

func getFunc(a func() int) int {
	return a()
}

func main() {
	a := getFunc(func() int {
		return int(math.Pow(5, 5))
	})
	fmt.Println(a)

}
