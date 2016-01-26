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

	db_host := os.Getenv("DATABASE_HOST")
	if db_host == "" {
		fmt.Printf("No db host given, falling back to localhost\n")
		db_host = "localhost"
	} else {
		fmt.Printf("Connecting to %v\n", db_host)
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
