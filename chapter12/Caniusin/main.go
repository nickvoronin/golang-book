package main

import (
	"fmt"
	"golang-book/chapter11/Caniusin/math"
)


func main() {
	number := []float64{34, 54, 1}
	avg := math.Average(number)
	fmt.Println(avg)
}