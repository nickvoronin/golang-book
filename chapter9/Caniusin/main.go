package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{2,2,4}
	c1 := Circle{2,2,3}
	r := Rectangle{2,3}
	fmt.Println(totalArea(&c, &c1, &r))
	fmt.Println(totalPerimeter(&c, &c1, &r))
}

type Circle struct{
	x, y, r float64
}
type Rectangle struct {
	x,y float64
}
func (c *Circle) area() float64{
	return math.Pi * c.r*c.r
}
func (r *Rectangle) area() float64{
	return  r.y *r.x
}
func (r *Rectangle) perimeter() float64{
	return 2*(r.x+r.y)
}
func (c *Circle) perimeter() float64{
	return 2* math.Pi * c.r
}
type Shape interface {
	area() float64
	perimeter() float64
}
func totalArea(shapes ...Shape) float64  {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func totalPerimeter(shapes ...Shape) float64  {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}