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

	sbyte := []byte(`{"ShipID":1,"ShipClass":"Fighter","Captain":{"name":"Jaro","clearace level":10,"access codes":["ADA","LAL"]}}`)

	si := new(ShipInfo)

	err := json.Unmarshal(sbyte, &si)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(si.ShipID, si.ShipClass, si.Captain.Name, si.Captain.SecurityClearance)

	m := make(map[string]int)
	sbyte = []byte(`{"Item1":1,"Item2":2}`)
	err = json.Unmarshal(sbyte, &m)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(m)

	s := []int{}
	sbyte = []byte(`[1,2,3,4,5]`)
	err = json.Unmarshal(sbyte, &s)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(s)

}
