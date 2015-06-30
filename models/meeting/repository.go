package meeting

import (
	"github.com/maxwellhealth/bongo"
	"gopkg.in/mgo.v2/bson"
)

var CollectionName string = "meetings"

func GetAll(dbConnection *bongo.Connection) []Meeting {
	results := dbConnection.Collection(CollectionName).Find(nil)
	m := &Meeting{}
	ms := make([]Meeting, 0)

	for results.Next(m) {
		ms = append(ms, *m)
	}
	return ms
}

func GetOne(id string, dbConnection *bongo.Connection) (*Meeting, error) {
	meeting := &Meeting{}
	err := dbConnection.Collection(CollectionName).FindById(bson.ObjectIdHex(id), meeting)

	return meeting, err
}

func Delete(id string, dbConnection *bongo.Connection) error {
	m := &Meeting{}

	m, err := GetOne(id, dbConnection)
	if err == nil {
		err = dbConnection.Collection(CollectionName).DeleteDocument(m)
	}
	return err
}
