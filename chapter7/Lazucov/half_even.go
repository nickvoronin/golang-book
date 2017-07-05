package main

import "fmt"

// мне кажется что в задании ошибка "half(1) должна вернуть (0, false),
// в то время как half(2) вернет (1, true)." 2/1 = 1 (нечет false)

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
