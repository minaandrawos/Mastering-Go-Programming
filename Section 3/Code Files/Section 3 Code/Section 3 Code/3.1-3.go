package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s1:", s1)
	s2 := make([]int, 2)
	n := copy(s2, s1[2:4])
	fmt.Println("Number of items copied: ", n)
	s2[0] = 10
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
}
