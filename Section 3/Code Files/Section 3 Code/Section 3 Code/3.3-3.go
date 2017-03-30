package main

import (
	"fmt"
	"time"
)

func main() {
	word := "Hello"

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(word)
	}()
	fmt.Println("Hello")
	word = "World"
	time.Sleep(3 * time.Second)
}
