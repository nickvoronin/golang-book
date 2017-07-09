package main

import "fmt"
import "golang-book/chapter11/Voronin/math"

func main() {
	xs := []float64{1,2,3,4}
	fmt.Println(math.Average(xs))
	fmt.Println(math.Max(xs))
	fmt.Println(math.Min(xs))
}