package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type crewMember struct {
	ID           int    `bson: "id"`
	Name         string `bson:"name"`
	SecClearance int    `bson:"security clearance"`
	Position     string `bson: "position"`
}

type Crew []crewMember

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//get collection
	personnel := session.DB("Hydra").C("Personnel")

	//Get number of documents in the collection
	n, _ := personnel.Count()
	log.Println("Number of personnel is ", n)

	//Perform simple query
	cm := crewMember{}
	personnel.Find(bson.M{"id": 3}).One(&cm)
	log.Println(cm)

	//Query with expression

	query := bson.M{
		"security clearance": bson.M{
			"$gt": 3,
		},
		"position": bson.M{
			"$in": []string{"Mechanic", "Biologist"},
		},
	}

	var crew Crew
	log.Println(query)
	err = personnel.Find(query).All(&crew)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Query results: ", crew)

	//Use select to get names only

	//names is of type []struct{Name string}
	names := []struct {
		Name string `bson:"name"`
	}{}

	/*
		equivalent to:

	*/

	err = personnel.Find(query).Select(bson.M{"name": 1}).All(&names)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(names)

	/*
				//Insert
				newcr := crewMember{ID: 18, Name: "Kaya Gal", SecClearance: 4, Position: "Biologist"}
				if err := personnel.Insert(newcr); err != nil {
					log.Fatal(err)
				}

			//update
			err = personnel.Update(bson.M{"id": 16}, bson.M{"$set": bson.M{"position": "Engineer III"}})
			if err != nil {
				log.Fatal(err)
			}

		//remove
		if err := personnel.Remove(bson.M{"id": 18}); err != nil {
			log.Fatal(err)
		}
	*/

	//Concurrent access
	var wg sync.WaitGroup
	count, _ := personnel.Count()
	wg.Add(count)
	for i := 1; i <= count; i++ {
		go readId(i, session.Copy(), &wg)
	}
	wg.Wait()

	//CSVToMongo(session.DB("Hydra").C("Personnel"))
}

func readId(id int, sessionCopy *mgo.Session, wg *sync.WaitGroup) {
	defer func() {
		sessionCopy.Close()
		wg.Done()
	}()
	p := sessionCopy.DB("Hydra").C("Personnel")
	cm := crewMember{}
	err := p.Find(bson.M{"id": id}).One(&cm)
	if err != nil {
		log.Printf("Could not retrieve id %d, error %s \n", id, err.Error())
		return
	}
	log.Println(cm)
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
