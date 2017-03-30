package main

import (
	"fmt"
)

func main() {
	helloPrinted := make(chan bool)

	go waitAndSay(helloPrinted, "world")

	fmt.Println("Hello")

	helloPrinted <- true

	if b := <-helloPrinted; b {
		fmt.Println("Program now signalled to exit")
	}
}

func waitAndSay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}
	c <- true
}
