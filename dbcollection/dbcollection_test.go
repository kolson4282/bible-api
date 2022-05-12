package dbcollection_test

import (
	"testing"

	"github.com/kolson4282/tdd-bible-api/dbcollection"
)

var DBVARS = dbcollection.DBVars{
	DB_USER:     "user",
	DB_PASSWORD: "dbpass",
	DB_HOST:     "localhost:3305",
	DB_NAME:     "bible_api",
}

func TestCharactersTable(t *testing.T) {
	dc := dbcollection.NewDBCollection(DBVARS)
	characters, _ := dc.GetCharacters()
	if len(characters) == 0 {
		t.Fatalf("No Characters Returned")
	}
	if characters[0].ID != 1 {
		t.Errorf("Wanted character ID 1, got %d", characters[0].ID)
	}
	for i, char := range characters {
		if char.Name == "" {
			t.Errorf("Invalid character name at postion %d", i)
		}
	}
}
