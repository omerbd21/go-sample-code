package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("5")
		wg.Done()
	}()
	go func() {
		fmt.Println("Am I fatser than 5?")
		wg.Done()
	}()
	wg.Wait()
}
