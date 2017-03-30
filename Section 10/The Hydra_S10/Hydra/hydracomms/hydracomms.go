package main

import (
	"Hydra/hydracomms/hydramessages"
	//hydraproto "Hydra/hydracomms/hydramessages/protobuff"
	"flag"
	"strings"
	"log"
	//"Hydra/hydracomms/hydramessages/thrift/gen-go/hydraThrift"
	"Hydra/hydracomms/hydramessages/gob"
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
	/*
	ship := &hydraproto.Ship{
		Shipname:    "Hydra",
		CaptainName: "Jala",
		Crew: []*hydraproto.Ship_CrewMember{
			&hydraproto.Ship_CrewMember{1, "Kevin", 5, "Pilot"},
			&hydraproto.Ship_CrewMember{2, "Jade", 4, "Tech"},
			&hydraproto.Ship_CrewMember{3, "Wally", 3, "Enginneer"},
		},
	}


	ship := &hydraThrift.Ship{
		Shipname:    "Hydra",
		CaptainName: "Jala",
		Crew: []*hydraThrift.CrewMember{
			&hydraThrift.CrewMember{1, "Kevin", 5, "Pilot"},
			&hydraThrift.CrewMember{2, "Jade", 4, "Tech"},
			&hydraThrift.CrewMember{3, "Wally", 3, "Enginneer"},
		},
	}
	*/

	ship := &hydragob.Ship{
		Shipname:    "Hydra",
		CaptainName: "Jala",
		Crew: []hydragob.CrewMember{
			hydragob.CrewMember{1, "Kevin", 5, "Pilot"},
			hydragob.CrewMember{2, "Jade", 4, "Tech"},
			hydragob.CrewMember{3, "Wally", 3, "Enginneer"},
		},
	}

	if err := hydramessages.EncodeAndSend(hydramessages.GOB, ship, address) ; err != nil {
		log.Println(err)
	}
}

func runServer(address string) {
	for ship := range hydramessages.ListenAndDecode(hydramessages.GOB, address) {
		log.Println(ship)
	}
}
