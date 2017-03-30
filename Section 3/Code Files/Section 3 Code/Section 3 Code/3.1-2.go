package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[2:4]
	c := a[:3]
	d := a[3:]

	fmt.Println("Slices a:", a, " b:", b, " c:", c, " d:", d)

	fmt.Println("What a actually sees: ", b[:cap(b)])

	fmt.Println("b slice only:", b[:len(b)])

}
