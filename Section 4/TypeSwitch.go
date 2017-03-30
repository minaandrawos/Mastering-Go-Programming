package main

import "fmt"

func main() {
	printType("text")
	printType(3)
	printType(4.0)
}

func printType(i interface{}) {
	switch i := i.(type) {
	case string:
		fmt.Println("This is a string type", i)
	case int:
		fmt.Println("This is an int type", i)
	case float32:
		fmt.Println("This is a float type", i)
	}
}
