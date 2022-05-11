package dbcollection_test

import (
	"testing"

	"github.com/kolson4282/tdd-bible-api/dbcollection"
)

func TestCharactersTable(t *testing.T) {
	dc := dbcollection.NewDBCollection()
	characters, _ := dc.GetCharacters()
	if len(characters) == 0 {
		t.Fatalf("No Characters Returned")
	}
	if characters[0].ID != 1 {
		t.Errorf("Wanted character ID 1, got %d", characters[0].ID)
	}
}
