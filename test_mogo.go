package main

import (
	"fmt"
	"log"
	"versionHistoryService/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"versionHistoryService/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main_() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "asdf"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
