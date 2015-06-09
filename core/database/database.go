package database

import (
	"gopkg.in/mgo.v2"
	"log"
)

type Database struct {
	session *mgo.Session
	dbName string
}

func Connect(hostAddress string, dbName string) *Database {
	session, err := mgo.Dial(hostAddress)
	if err != nil {
		panic(err)
	}

	return &Database{session: session, dbName: dbName}
}

func (d *Database) Close() {
	d.session.Close()
}

func (d *Database) Insert(model Model) {
	err := d.session.DB(d.dbName).C(model.GetCollectionName()).Insert(model)
	if err != nil {
		log.Fatal(err)
	}
}
