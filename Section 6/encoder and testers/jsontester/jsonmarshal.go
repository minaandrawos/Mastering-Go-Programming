package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type CrewMember struct {
		ID                int      `json:"id,omitempty"`
		Name              string   `json:"name"`
		SecurityClearance int      `json:"clearace level"`
		AccessCodes       []string `json:"access codes"`
	}

	type ShipInfo struct {
		ShipID    int
		ShipClass string
		Captain   CrewMember
	}

	cm := CrewMember{Name: "Jaro", SecurityClearance: 10, AccessCodes: []string{"ADA", "LAL"}}

	si := ShipInfo{1, "Fighter", cm}

	b, err := json.Marshal(&si)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	m := map[string]int{"Item1": 1, "Item2": 2}
	b, err = json.Marshal(&m)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	s := []int{1, 2, 3, 4}
	b, err = json.Marshal(&s)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

}
