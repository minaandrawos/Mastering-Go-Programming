package main

import (
	"encoding/xml"
	"fmt"
)

func main() {

	type CrewMember struct {
		XMLName           xml.Name `xml:"member"`
		ID                int      `xml:"id,omitempty"`
		Name              string   `xml:"name,attr"`
		SecurityClearance int      `xml:"clearace,attr"`
		AccessCodes       []string `xml:"codes>code"`
	}

	type ShipInfo struct {
		XMLName   xml.Name `xml:"ship"`
		ShipID    int      `xml:"ShipDetails>ShipID"`
		ShipClass string   `xml:"ShipDetails>ShipClass"`
		Captain   CrewMember
	}

	cm := CrewMember{Name: "Jaro", SecurityClearance: 10, AccessCodes: []string{"ADA", "LAL"}}

	si := ShipInfo{ShipID: 1, ShipClass: "Fighter", Captain: cm}

	b, err := xml.MarshalIndent(&si, " ", "	")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	m := map[string]int{"item1": 1, "item2": 2}
	b, err = xml.Marshal(m)
	fmt.Println(string(b))

	s := []int{1, 2, 3, 4}
	b, err = xml.Marshal(&s)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

}
