package main

import "fmt"

type CMember struct {
	name      string
	age       int
	address   string
	rank      string
	clearance int
}

func main() {
	//Declare a struct

	//cm := CMember{"Kevin",32,"Station Mars","Sergeant",5}

	//or

	/*
	  cm := CMember{
	    name: "Kevin",
	    clearance: 5,
	    age: 32,
	    rank: "Sergeant",
	    address: "Station Mars",
	  }*/

	//or

	var cm CMember
	cm.name = "Kevin"
	cm.age = 32
	cm.address = "Station Mars"
	cm.rank = "Sergeant"
	cm.clearance = 5

	cmp := &cm
	cmp.clearance = 4

	var crew []CMember
	crew = append(crew, cm, CMember{"Jo", 32, "Station Mars", "Sergeant", 5})

	for i, v := range crew {
		fmt.Println(i, v)
	}

	//maps with structs
	//var m map[string]CMember
	//m = make(map[string]CMember)
	//m["Keving"] = cm

	//or

	m := map[string]CMember{
		"Kevin": CMember{name: "Kevin", address: "Stations Mars"},
		"Jo":    CMember{name: "Jo", address: "Station Jupiter"},
	}

	//add
	m["Cisco"] = CMember{name: "Cisco", address: "Station Mars", clearance: 5}
	//retrieve
	elem := m["Jo"]
	fmt.Println(elem)
	//Check if the value exists
	v, ok := m["Jo"]
	//delete
	delete(m, "Jo")

	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}
}

func (cm CMember) PrintSecurityClearance() {
	fmt.Println(cm.clearance)
}
