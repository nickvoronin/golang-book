package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x, y float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.x + r.y)
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{3, 4}

	fmt.Println(r.perimeter())
	fmt.Println(c.perimeter())
}
