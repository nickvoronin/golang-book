package main

import "fmt"

func halfNumber(num int) bool {
	return num%2 == 0
}

func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)

	if halfNumber(input) {
		fmt.Println("1", "true")
	} else {
		fmt.Println("0", "false")
	}
}
