package main

import "fmt"

type level int

func main() {
	
	var sl level = 10
	//sl := new(level)
	
	sl.raiseShieldLevel(4)
	sl.raiseShieldLevel(5)
	
	//fmt.Println(*sl)
	fmt.Println(sl)
	
}

func (lv *level)raiseShieldLevel(i int){

	if *lv==0 {
		*lv = 1
	}
	
	*lv = (*lv) * level(i)
	
}