package main

import "fmt"

func swap(xPtr *int, yPtr *int) {
	var temp int = *yPtr
	*yPtr = *xPtr
	*xPtr = temp
	return
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}
