package main

import "fmt"

func main() {
	n := 1
	//A for loop
	for i := 10; i > 0; i-- {
		n *= i
	}
	fmt.Println("Result: ", n)

	n = 1
	//A Go While
	for n <= 50 {
		n += n
	}
	fmt.Println("Result: ", n)
}
