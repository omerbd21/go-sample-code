package main

import "fmt"

func main() {
	s := struct {
		name     string
		nickname string
	}{
		name:     "benda",
		nickname: "bendibenda",
	}
	fmt.Println(s.name, s.nickname)
}
