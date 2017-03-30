package main

import (
	"encoding/xml"
	"log"
	"os"
)

type CrewMember struct {
	XMLName           xml.Name `xml:"member"`
	ID                int      `xml:"id,omitempty"`
	Name              string   `xml:"name,attr"`
	SecurityClearance int      `xml:"clearace,attr"`
	AccessCodes       []string `xml:"codes>code"`
}

type ShipInfo struct {
	XMLName   xml.Name `xml:"ship"`
	ShipID    int      `xml:"ShipInfo>ShipID"`
	ShipClass string   `xml:"ShipInfo>ShipClass"`
	Captain   CrewMember
}

func main() {
	file, err := os.Create("xmlfile.xml")
	if err != nil {
		log.Fatal("Could not create file", err)
	}
	defer file.Close()
	//fill some data
	cm := CrewMember{Name: "Jaro", SecurityClearance: 10, AccessCodes: []string{"ADA", "LAL"}}
	si := ShipInfo{ShipID: 1, ShipClass: "Fighter", Captain: cm}

	enc := xml.NewEncoder(file)
	enc.Indent(" ", "  ")
	enc.Encode(si)
	if err != nil {
		log.Fatal("Could not encode xml file", err)
	}
}
