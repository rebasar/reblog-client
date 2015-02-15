package main

import (
	"gopkg.in/mgo.v2"
)

func uploadEntry(entry Entry) {
	session, err := mgo.Dial(Server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB(Database).C("entries")
	err = c.Insert(&entry)
	if err != nil {
		panic(err)
	}
}
