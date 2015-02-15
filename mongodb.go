package main

import (
	"gopkg.in/mgo.v2"
)

func uploadEntry(entry Entry) {
	session, err := mgo.Dial(Server)
	checkError(err)
	defer session.Close()
	c := session.DB(Database).C("entries")
	err = c.Insert(&entry)
	checkError(err)
}
