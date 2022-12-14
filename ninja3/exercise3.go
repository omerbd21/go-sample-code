package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	birthYear := 2000
	for t.Year() > birthYear {
		fmt.Println(birthYear)
	}
}
