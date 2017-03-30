package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s1 = append(s1, 4, 5, 6)
	s2 := []int{7, 8, 9}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
