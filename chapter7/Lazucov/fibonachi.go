package main

import "fmt"

func fb(n uint) uint {
	if n <= 1 {
		return n
	} else {
		return fb(n-1) + fb(n-2)
	}

}

func main() {
	fmt.Println(fb(20))
}
