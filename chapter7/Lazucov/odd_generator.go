package main

import "fmt"

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

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 0
	fmt.Println(nextOdd()) // 2
	fmt.Println(nextOdd()) // 4
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
