package main

import (
	"fmt"
	"net/http"
	"os"
	"versionHistoryService/Godeps/_workspace/src/gopkg.in/mgo.v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db_host := os.Getenv("MONGOLAB_URI")
	if db_host == "" {
		db_host = "localhost"
	}

	session, err := mgo.Dial(db_host)
	if err != nil {
		panic(err)
	}

	db := session.DB("av_version")

	defer session.Close()

	r := Router(db)

	fmt.Println("Starting server on :" + port)
	http.ListenAndServe(":"+port, r)
}
