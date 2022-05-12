package dbcollection_test

//dbcollection_test needs to be run with a clean (empty) database

import (
	"testing"

	"github.com/kolson4282/tdd-bible-api/dbcollection"
	"github.com/kolson4282/tdd-bible-api/graph/model"
)

var DBVARS = dbcollection.DBVars{
	DB_USER:     "user",
	DB_PASSWORD: "dbpass",
	DB_HOST:     "localhost:3305",
	DB_NAME:     "bible_api",
}

func TestCharactersTable(t *testing.T) {
	dc := dbcollection.NewDBCollection(DBVARS)

	t.Run("Get All Characters", func(t *testing.T) {
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
	})

	t.Run("Create New Character", func(t *testing.T) {
		newCharacter := model.NewCharacter{
			Name:        "Adam",
			Description: "First Man",
		}
		character, _ := dc.CreateCharacter(newCharacter)
		if character.Name != newCharacter.Name {
			t.Errorf("Character created incorrectly, wanted %s, got %s",
				newCharacter.Name, character.Name)
		}

		characters, _ := dc.GetCharacters()
		found := false
		for _, char := range characters {
			if char.Name == newCharacter.Name && char.Description == newCharacter.Description {
				found = true
			}
		}
		if !found {
			t.Error("Character not added to list of all characters")
		}
	})

	t.Run("Can't create new Character if character already exists", func(t *testing.T) {
		newCharacter := model.NewCharacter{
			Name:        "Adam",
			Description: "First Man",
		}
		_, err := dc.CreateCharacter(newCharacter)

		if err.Error() != "character already exists" {
			t.Errorf("tried to create duplicate character")
		}
		characters, _ := dc.GetCharacters()
		found := false
		for _, char := range characters {
			if char.Name == newCharacter.Name && char.Description == newCharacter.Description {
				found = true
			}
		}
		if !found {
			t.Error("Character was not found in list of all characters")
		}
	})
}
