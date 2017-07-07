package main

import "fmt"

func fibonacci(number int) int {
	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	return fibonacci(number-1) + fibonacci(number-2)

}

func main() {
	for i := 0; i < 7; i++ {
		fmt.Println(fibonacci(i))
	}
}
