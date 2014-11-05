package api

import (
	"testing"
)

func TestAlbumPut(t *testing.T) {
	// This is a successful PUT
	jsonData := `{
			     	"name": "The Earth Is Not a Cold Dead Place",
					"year": "2003",
					"genre": "post-rock",
					"artist": "Explosions in the Sky"
				 }`
	result := PutAlbum(jsonData)
	if result != "" {
		t.Error("Expected \"\", got ", result)
	}

}
