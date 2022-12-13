package main

import "fmt"

func getFunc() func() {
	return func() {
		fmt.Println("Inside the machine")
	}
}

func main() {
	a := getFunc()
	a()

}
