package hydradblayer

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoDataStore struct {
	*mgo.Session
}

func NewMongoStore(conn string) (*mongoDataStore, error) {
	log.Println(conn)
	session, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}
	return &mongoDataStore{Session: session}, nil
}

//In case of mongodb, the id field doesn't auto increment as the case was with MySQL.
//So the json string used in the API post request body need to supply the id.
func (ms *mongoDataStore) AddMember(cm *CrewMember) error {
	session := ms.Copy()
	defer session.Close()
	personnel := session.DB("Hydra").C("Personnel")
	err := personnel.Insert(cm)
	return err
}

func (ms *mongoDataStore) FindMember(id int) (CrewMember, error) {
	session := ms.Copy()
	defer session.Close()
	personnel := session.DB("Hydra").C("Personnel")
	cm := CrewMember{}
	err := personnel.Find(bson.M{"id": id}).One(&cm)
	return cm, err
}

func (ms *mongoDataStore) AllMembers() (crew, error) {
	session := ms.Copy()
	defer session.Close()
	personnel := session.DB("Hydra").C("Personnel")
	members := crew{}
	err := personnel.Find(nil).All(&members)
	return members, err
}
