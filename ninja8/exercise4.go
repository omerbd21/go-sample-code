package main

import (
	"fmt"
	"sort"
)

type user struct {
	First string
	Age   int
}

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2}
	xs := []string{"random", "rainbow", "sahar", "omer", "meitar"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}
