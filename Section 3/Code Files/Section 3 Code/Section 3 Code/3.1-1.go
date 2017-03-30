package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s1:", s1)
	s2 := s1[2:4]
	s2[0] = 10
	fmt.Println("S1: ", s1)

}
