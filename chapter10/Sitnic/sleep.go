package main

import (
	"fmt"
	"time"
)

func Sleep(s int) {
	<-time.After(time.Second * time.Duration(s))
}

func main() {
	i := 1
	for i <= 10 {
		Sleep(1)
		fmt.Println(i)
		i = i + 1
	}
}
