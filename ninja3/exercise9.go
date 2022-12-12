package main

import (
	"fmt"
)

func main() {
	favSport := "Football"
	switch favSport {
	case "Football":
		fmt.Println("YESSSSS")
	case "Basketball":
		fmt.Println("What are you from Yugoslavia?")
	case "Soccer":
		fmt.Println("For the last time, it's football!!")
	}
}
