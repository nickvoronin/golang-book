package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 1
	y := 2
	fmt.Println("before x =", x, "y =", y)
	swap(&x, &y)
	fmt.Println("after x =", x, "y =", y)
}
