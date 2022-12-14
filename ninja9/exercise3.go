package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	x := 0
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		x++
		runtime.Gosched()
		fmt.Println(x)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(x)
}
