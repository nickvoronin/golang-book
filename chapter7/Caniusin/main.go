
package main

import "fmt"


func main() {
	//sum
	numbers := []int{1,2,3}
	fmt.Println("Sum is: ", sum(numbers))

	//half
	n := 2
	fmt.Print("Half:  ")
	fmt.Println(half(n))

	//max in slice
	fmt.Println("Max number is:", max(34, 55, 0.001 , 222.001,222.0001, 5, 56.54 ))

	//odd generator
	odd := makeOddGenerator()
	i := 0
	fmt.Println("Odd Numbers:  ")
	for i<5{
		fmt.Println(odd())
		i++
	}



	//fibonaci
	number := 6
	fmt.Println("Fibonaci:",fib(number))

}

//sum
func sum(numbers []int) int {
	var result int
	for  _, num := range numbers{
		result += num
	}
	return result
}

//Half
func half(x int) (int, bool) {

	return x/2, x%2 == 0
}

//max in slice
func max(num ...float64) float64 {
	var res float64
	for _, n := range num {
		if n > res {
			res = n
		}
	}
	return res
}

//oddgenerator
func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		if i%2 > 0 {
			ret = i
		} else {
			i += 1
			ret = i
		}
		i += 1
		return
	}
}

//Fibonaci
func fib(n int)  int{
	if n < 2{
		return n
	}
	return fib(n-1) + fib(n-2)
}