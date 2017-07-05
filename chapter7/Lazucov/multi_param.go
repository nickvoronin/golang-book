package main

import "fmt"

func multi(args ...float64) float64 {
	result := 0.0
	for _, v := range args {
		if v > result {
			result = v
		}
	}
	return result
}
func main() {
	fmt.Println(multi(3, 2, 17, 1231, 12.34, 1.2441211))
}
