package main

import "fmt"

func main() {
	x := make([]int, 10, 10)
	x[0] = 42
	x[1] = 43
	x[2] = 44
	x[3] = 45
	x[4] = 46
	x[5] = 47
	x[6] = 48
	x[7] = 49
	x[8] = 50
	x[9] = 51
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}
