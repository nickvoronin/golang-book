package main

import "fmt"

import "golang-book/chapter11/Sitnic/math"

func main() {

	x := []float64{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println(min(x...))
	fmt.Println(max(x...))
}
