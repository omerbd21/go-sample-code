package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	x := 0
	var wg sync.WaitGroup
	var mx sync.Mutex
	wg.Add(1)
	go func() {
		mx.Lock()
		x++
		runtime.Gosched()
		mx.Unlock()
		fmt.Println(x)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(x)
}
