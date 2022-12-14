package canine

import "strings"

func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
