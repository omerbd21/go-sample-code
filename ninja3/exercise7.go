package main

import (
	"fmt"
)

func main() {
	x := "Yalla Haifa"
	if x == "Yalla Haifa" {
		fmt.Println(x)
	} else if x == "Yalla Maccabi" {
		fmt.Println(x)
	} else {
		fmt.Println("Rak Behaifa Yesh Maccabi - R.B.Y.M")
	}
}
