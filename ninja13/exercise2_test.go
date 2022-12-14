package canine

import "testing"

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("one")
	}
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "wanted", 3)
	}
}
