package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1) // buffered with size 1

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Anonymous function writing to c1"
	}()

	// timeout before the above funciton could write to c1
	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Anonymous function writing to c2"
	}()

	// Won't timeout because channel write was faster
	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
