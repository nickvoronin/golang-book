package main

import "fmt"

func main() {
	fmt.Print("Enter a number in feets: ")
	var feets float64
	fmt.Scanf("%f", &feets)

	meters := feets * 0.3048

	fmt.Println(meters)
}
