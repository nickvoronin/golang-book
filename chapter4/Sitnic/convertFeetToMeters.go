package main

import "fmt"

func main() {

	fmt.Print("Enter a number to convert feet to meters: ")

	var feet float64

	fmt.Scanf("%f", &feet)

	var meters = (feet / 3.2808)

	fmt.Println(meters)
}
