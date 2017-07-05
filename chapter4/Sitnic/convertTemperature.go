package main

import "fmt"

func main() {

	fmt.Print("Enter a number to convert temperature from fahrenheit to celsius: ")

	var fahrenheit float64

	fmt.Scanf("%f", &fahrenheit)

	var celsius = (fahrenheit - 32) * 5 / 9

	fmt.Println(celsius)
}
