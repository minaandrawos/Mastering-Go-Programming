package main

import (
	"Hydra/hydracomms/hydramessages"
	hydraproto "Hydra/hydracomms/hydramessages/protobuff"
	"flag"
	"strings"
	"log"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8080", "address? host:port ")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runClient(address string) {
	ship := &hydraproto.Ship{
		Shipname:    "Hydra",
		CaptainName: "Jala",
		Crew: []*hydraproto.Ship_CrewMember{
			&hydraproto.Ship_CrewMember{1, "Kevin", 5, "Pilot"},
			&hydraproto.Ship_CrewMember{2, "Jade", 4, "Tech"},
			&hydraproto.Ship_CrewMember{3, "Wally", 3, "Enginneer"},
		},
	}

	hydramessages.EncodeAndSend(hydramessages.Protobuf, ship, address)
}

func runServer(address string) {
	for ship := range hydramessages.ListenAndDecode(hydramessages.Protobuf, address) {
		log.Println(ship)
	}
}
