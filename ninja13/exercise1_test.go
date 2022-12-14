package canine

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	c := canineAge(canine{name: "a", age: 5})
	if c != 5 {
		t.Error("got", c, "wanted", 5)
	}
}

func ExampleMain() {
	c := canine{
		name: "a",
		age:  5,
	}
	fmt.Println(canineAge(c))
	// Output:
	// 5
}

func BenchmarkMain(b *testing.B) {
	c := canine{
		name: "a",
		age:  5,
	}
	for i := 0; i < b.N; i++ {
		canineAge(c)
	}
}
