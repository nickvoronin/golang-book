package main

import "fmt"

func half(x int) (int, bool) {
	y := x / 2
	if y%2 > 0 {
		return y, false
	}
	return y, true
}

func main() {
	fmt.Println(half(2))
}
