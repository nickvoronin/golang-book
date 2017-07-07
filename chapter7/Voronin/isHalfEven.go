package main

import "fmt"

func isHalfEven(number int) bool {
	return number%2 == 0
}

func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)

	if isHalfEven(input) {
		fmt.Println("The half of", input, "is even")
	} else {
		fmt.Println("The half of", input, "is not even")
	}
}
