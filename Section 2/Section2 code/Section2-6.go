package main

import "fmt"

type Set map[string]struct{}

func main() {
	s := make(Set)
	//add items
	s["Item1"] = struct{}{}
	s["Item2"] = struct{}{}
	//get and print items
	fmt.Println(getSetValues(s))
}

func getSetValues(s Set) []string {
	var retVal []string
	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}
