package main

import "fmt"

func main() {
	subslice := testSubSlice()
	fmt.Println(subslice, " length: ", "remaining underlying array: ", subslice[:cap(subslice)])
}

func testSubSlice() []int {
	s := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	return s[1:4]
}
