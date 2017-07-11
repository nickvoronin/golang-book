package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("Swap numbers are:" ,x, y)
}