package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First`
	Last    string   `json:"Last`
	Age     int      `json:"Age`
	Sayings []string `json:"Sayings`
}

func main() {
	j := `[{"First": "James", "Last": "Bond", "Age": 32, "Sayings":["Shaken, not stirred", "Youth is no good"]}]`
	var people []person
	if err := json.Unmarshal([]byte(j), &people); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people)

}
