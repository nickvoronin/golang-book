package main

import (
	"fmt"
	"time"
)

func Sleep(s float32) {
	<-time.After(time.Second * time.Duration(s))
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			Sleep(0.1)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			Sleep(0.01)
		}
	}()

	go func() {
		for {
			select {
				case msg1 := <- c1:
					fmt.Println("Message 1", msg1)
				case msg2 := <- c2:
					fmt.Println("Message 2", msg2)
				case <- time.After(time.Second):
					fmt.Println("timeout")
				default:
					fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}