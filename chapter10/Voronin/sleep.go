package main

import ("fmt"; "time")

func Sleep(seconds int) {
	<- time.After(time.Second * time.Duration(seconds))
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)


	go func() {
		for {
			c1 <- "from 1"
			Sleep(2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			Sleep(3)
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
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}