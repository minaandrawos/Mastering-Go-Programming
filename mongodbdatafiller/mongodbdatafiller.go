package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"flag"

	mgo "gopkg.in/mgo.v2"
)

type crewMember struct {
	ID           int    `bson:"id"`
	Name         string `bson:"name"`
	SecClearance int    `bson:"security clearance"`
	Position     string `bson:"position"`
}

type Crew []crewMember

func main() {
	mgoaddress := flag.String("a", "mongodb://127.0.0.1", "Mongodb connection address")
	flag.Parse()
	//session, err := mgo.Dial("localhost")
	session, err := mgo.Dial(*mgoaddress)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//get collection
	personnel := session.DB("Hydra").C("Personnel2")
	CSVToMongo(personnel)
}

func CSVToMongo(c *mgo.Collection) {
	file, err := os.Open("Crews.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comment = '#'
	var crew []interface{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		if len(record) != 4 {
			continue
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Println(err)
			continue
		}
		sc, err := strconv.Atoi(record[2])
		if err != nil {
			log.Println(err)
			continue
		}
		crew = append(crew, crewMember{id, record[1], sc, record[3]})
	}
	log.Println("Crew found in CSV: ", crew)
	err = c.Insert(crew...)
	if err != nil {
		log.Fatal(err)
	}
}
