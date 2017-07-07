package main

import "fmt"
import "golang-book/chapter11/Lazucov/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	min := math.Min(xs)
	max := math.Max(xs)
	fmt.Println(min)
	fmt.Println(max)
	fmt.Println(avg)
}
