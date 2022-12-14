package main

import (
	"fmt"
	"sync"
)

func main() {
	x := 0
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		x++
		fmt.Println(x)
		wg.Done()
	}()
	wg.Done()
	fmt.Println(x)

}
