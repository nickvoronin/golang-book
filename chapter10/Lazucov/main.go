package main

import (
	"fmt"
	"time"
)

func Sleep(s int) {
	<-time.After(time.Second * time.Duration(s))
}

func main() {
	for {
		Sleep(10)
		fmt.Println("hello")
	}
}
