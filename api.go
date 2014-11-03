package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmcvetta/neoism"
)

type Album struct {
	Name string
	Year string
	Recs string
}

// AlbumSelect represents the data passed up to select
// an album from the database
type AlbumSelect struct {
	Name   string
	Year   string
	Artist string
	Genre  string
}

/*
 * GetAlbum will return a list of albums that match
 * the given data, the data may select based on:
 * DATE, ARTIST, YEAR, and GENRE
 */
func GetAlbum(jsonData string) string {
	fmt.Println("GET ALBUM")

	// Pull selection data from json string
	var as AlbumSelect
	err := json.Unmarshal([]byte(jsonData), &as)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("FIELDS: %#v\n", as)

	// Pull Neo4j nodes from DB matching selecton params
	res := []struct {
		N string `json:"n.name"`
		Y string `json:"n.year"`
		R string `json:"n.recs"`
	}{}

	cq := neoism.CypherQuery{
		// We use regex matches (=~) to gracefully account for
		// missing fields, so we can use .*
		Statement: `
			MATCH (n:Album)
			WHERE n.name =~ {name} AND n.year =~ {year}
			RETURN n.name, n.year;
		`,
		// DefMatch is substituting .* for us when necessary
		Parameters: neoism.Props{"name": DefMatch(as.Name), "year": DefMatch(as.Year)},
		Result:     &res,
	}
	DBConn.Cypher(&cq)

	// Turn the list of Nodes into a list of Albums
	albums := make([]Album, 1)
	for _, el := range res {
		name := el.N
		year := el.Y
		recs := el.R

		a := Album{Name: name, Year: year, Recs: recs}
		albums = append(albums, a)
	}

	// Turn the list of albums into a json representation
	jsonReturn, err := json.Marshal(albums)
	if err != nil {
		fmt.Println(err)
	}

	return string(jsonReturn)
}

func PutAlbum(jsonData string) string {
	// TODO: Write a data verification method

	fmt.Println("Received Data: ", jsonData)

	// Parse the json data into an album struct
	var a Album
	err := json.Unmarshal([]byte(jsonData), &a)
	if err != nil {
		fmt.Println(err)
	}

	// Create a new node in Neo4j DB
	// TODO: Refactor this without manual cypher query
	res := []struct {
		N neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement:  "CREATE (n:Album {name: {name}, year: {year}, recs: {recs}}) RETURN n",
		Parameters: neoism.Props{"name": a.Name, "year": a.Year, "recs": a.Recs},
		Result:     res,
	}
	DBConn.Cypher(&cq)

	// TODO: Create relationships to artist, genre

	return ""
}
