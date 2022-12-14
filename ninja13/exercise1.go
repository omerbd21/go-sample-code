package canine

type canine struct {
	name string
	age  int
}

func canineAge(c canine) int {
	return c.age
}
