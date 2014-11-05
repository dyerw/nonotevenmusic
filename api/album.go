package api

import (
	"encoding/json"
	"fmt"
	"github.com/jmcvetta/neoism"
	"net/url"
	"time"
)

type Album struct {
	Name      string
	Year      string
	Submitted int32
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
 * Put album will add an album to the database and return
 * nothing if successful and a json object with an error message
 * if unsuccessful.
 */
func PutAlbum(jsonData string, db *neoism.Database) string {
	// TODO: Write a data verification method

	// Parse the json data into an album struct
	var a Album
	err := json.Unmarshal([]byte(jsonData), &a)
	if err != nil {
		return "{ \"err\": \"Unable to parse json request\" }"
		fmt.Println(err)
	}

	// Set the submitted date to the current time
	a.Submitted = int32(time.Now().Unix())
	fmt.Println(a.Submitted)

	// Create a new node in Neo4j DB
	res := []struct {
		N neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement:  "CREATE (n:Album {name: {name}, year: {year}, submitted: {submitted}}) RETURN n",
		Parameters: neoism.Props{"name": a.Name, "year": a.Year, "submitted": a.Submitted},
		Result:     res,
	}
	db.Cypher(&cq)

	// TODO: Create relationships to artist, genre

	return ""
}

/*
 * GetAlbum will return a list of albums that match
 * the given data, the data may select based on:
 * DATE, ARTIST, YEAR, and GENRE
 */
func GetAlbum(urlArgs url.Values, db *neoism.Database) string {
	// Pull selection data from url arguments
	as := AlbumSelect{Name: urlArgs.Get("name"), Year: urlArgs.Get("year"), Genre: urlArgs.Get("genre"), Artist: urlArgs.Get("artist")}

	// Pull Neo4j nodes from DB matching selecton params
	res := []struct {
		N string `json:"n.name"`
		Y string `json:"n.year"`
		S int32  `json:"n.submitted"`
	}{}

	cq := neoism.CypherQuery{
		// We use regex matches (=~) to gracefully account for
		// missing fields, so we can use .*
		Statement: `
			MATCH (n:Album)
			WHERE n.name =~ {name} AND n.year =~ {year}
			RETURN n.name, n.year, n.submitted;
		`,
		// DefMatch is substituting .* for us when necessary
		Parameters: neoism.Props{"name": DefMatch(as.Name), "year": DefMatch(as.Year)},
		Result:     &res,
	}
	db.Cypher(&cq)

	// Turn the list of Nodes into a list of Albums
	albums := make([]Album, 1)
	for _, el := range res {
		name := el.N
		year := el.Y
		submitted := el.S

		a := Album{Name: name, Year: year, Submitted: submitted}
		albums = append(albums, a)
	}

	// Turn the list of albums into a json representation
	jsonReturn, err := json.Marshal(albums)
	if err != nil {
		fmt.Println(err)
	}

	return string(jsonReturn)
}
