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
		ShipID    int      `xml:"ShipInfo>ShipID"`
		ShipClass string   `xml:"ShipInfo>ShipClass"`
		Captain   CrewMember
	}

	bytes := []byte(` <ship>
				<ShipInfo>
 					<ShipID>1</ShipID>
 					<ShipClass>Fighter</ShipClass>
 				</ShipInfo>
 				<member name="Jaro" clearace="10">
 					<codes>
 						<code>ADA</code>
 						<code>LAL</code>
 					</codes>
 				</member>
 			 </ship>`)

	si := ShipInfo{}

	err := xml.Unmarshal(bytes, &si)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(si)

}
