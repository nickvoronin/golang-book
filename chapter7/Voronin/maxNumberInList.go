package main

import "fmt"

func maxNumberInList(numbers ...float64) float64 {
	max := numbers[0]
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	fmt.Println(maxNumberInList(1, 0, 9, 5))
}
