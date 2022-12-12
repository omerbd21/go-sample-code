package main

import "fmt"

func main() {
	x := map[string][]string{
		"Benda": {"Football", "DevOps"},
		"Omer":  {"Basketball", "Kiurtush"},
		"Sahar": {"Video Games", "Google"},
	}
	x["Yuvalis"] = []string{"Eating the head"}
	delete(x, "Yuvalis")
	for k, v := range x {
		fmt.Println(k)
		for i, z := range v {
			fmt.Println(i, z)
		}
	}

}
