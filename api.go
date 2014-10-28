package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmcvetta/neoism"
)

type Album struct {
	Name string
	Year int
	Recs int
}

func GetAlbum(jsonData string) string {
	return ""
}

func PutAlbum(jsonData string) string {
	fmt.Println("Received Data: ", jsonData)

	// Parse the json data into an album struct
	var a Album
	err := json.Unmarshal([]byte(jsonData), &a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a.Name)
		fmt.Println(a.Year)
		fmt.Println(a.Recs)
	}

	// Create a new node in Neo4j DB
	res := []struct {
		N neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement:  "CREATE (n:Album {name: {name}, year: {year}, recs: {recs}}) RETURN n",
		Parameters: neoism.Props{"name": a.Name, "year": a.Year, "recs": a.Recs},
		Result:     res,
	}
	DBConn.Cypher(&cq)

	return ""
}
