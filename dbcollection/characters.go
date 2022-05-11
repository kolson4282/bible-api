package dbcollection

import "github.com/kolson4282/tdd-bible-api/graph/model"

func (dc *DBCollection) GetCharacters() ([]*model.Character, error) {
	var characters []*model.Character
	dc.DB.Find(&characters)
	return characters, nil
}
