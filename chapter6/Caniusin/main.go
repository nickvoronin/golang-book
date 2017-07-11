package main

import "fmt"

func main() {
	x := []int{
			48,96,86,68,
			57,1,63,70,
			37,34,83,27,
			19,97, 9,17,
		}
		fmt.Println("Min:",min(x))
}

func min(n []int)  int{
	var min int
	for _, val :=  range n{
		if min == 0{
			min = val
		}
		if val < min{
			min = val
		}
	}
	return min
}