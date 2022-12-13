package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (s square) info() {
	fmt.Println(s.area())
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c circle) info() {
	fmt.Println(c.area())
}

func main() {
	s := square{
		side: 5,
	}
	c := circle{
		radius: 3,
	}
	s.info()
	c.info()

}
