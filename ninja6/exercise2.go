package main

import "fmt"

func main() {
	fmt.Println(foo(1, 2, 3))
	fmt.Println(bar([]int{1, 2, 3}))
}

func foo(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func bar(a []int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}
