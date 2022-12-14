package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var x int64
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		atomic.AddInt64(&x, 1)
		fmt.Println(x)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(x)
}
