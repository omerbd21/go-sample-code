package canine

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	ca := CenteredAvg([]int{1, 2, 3, 4, 5})
	if ca != 3 {
		t.Error("Got", ca, "Expected", 3)
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < 100; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5, 6}))
	// Output:
	// 3.5
}
