package main

import "fmt"

func main() {
	fmt.Print("Enter temperature in fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	var celsius = (fahrenheit - 32) * 5 / 9
	fmt.Println("Temperature in Celsius is:",celsius)
}
