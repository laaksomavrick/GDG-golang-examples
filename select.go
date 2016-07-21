package main

import (
	"fmt"
	"time"
)

func main() {
	//create two channels
	c1 := make(chan string)
	c2 := make(chan string)

	//create two IIFEs sending msg to chan
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	//Take actions in relation to chan case
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
