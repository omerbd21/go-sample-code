package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	birthYear := 2000
	for t.Year() >= birthYear {
		if t.Year() <= birthYear {
			fmt.Println(birthYear)
			birthYear++
		} else {
			break
		}
	}
}
