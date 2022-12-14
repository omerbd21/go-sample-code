package main

import "fmt"

type customErr struct {
	something string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("this is some kind of error:%v", ce.something)
}

func main() {
	c1 := customErr{
		something: "fell miserably",
	}
	idan(c1)
}

func idan(e error) {
	fmt.Println("idan ran - ", e)
}
