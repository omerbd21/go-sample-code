package main

import (
	"fmt"
)

func fur() func() string {
	s := ""
	return func() string {
		s = "hiiiii"
		return s
	}
}

func main() {
	fmt.Println(fur()())

}
