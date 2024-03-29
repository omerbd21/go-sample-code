package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type ByAge []user

func (b ByAge) Len() int           { return len(b) }
func (b ByAge) Less(i, j int) bool { return b[i].Age < b[j].Age }
func (b ByAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type ByLast []user

func (b ByLast) Len() int           { return len([]user(b)) }
func (b ByLast) Less(i, j int) bool { return b[i].Last < b[j].Last }
func (b ByLast) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that fot you?",
			"I would really prefer to be a secret agent",
		},
	}
	u3 := user{
		First: "M",
		Last:  "Hmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't",
			"Dear God, whar has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	users := []user{u1, u2, u3}
	fmt.Println(users)
	sort.Sort(ByAge(users))
	fmt.Println(users)
	sort.Sort(ByLast(users))
	fmt.Println(users)
	sort.Strings(u1.Sayings)
	sort.Strings(u2.Sayings)
	sort.Strings(u3.Sayings)
	fmt.Println(users)

}
