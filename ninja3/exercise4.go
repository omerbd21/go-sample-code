package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	birthYear := 2000
	for {
		if t.Year() >= birthYear {
			fmt.Println(birthYear)
			birthYear++
		} else {
			break
		}
	}
}
