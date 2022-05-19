package dbcollection

import (
	"errors"

	"github.com/kolson4282/tdd-bible-api/graph/model"
)

func (dc *DBCollection) GetCharacters() ([]*model.Character, error) {
	var characters []*model.Character
	dc.DB.Find(&characters)
	return characters, nil
}

func (dc *DBCollection) CreateCharacter(newCharacter model.NewCharacter) (*model.Character, error) {
	character := model.Character{
		Name:        newCharacter.Name,
		Description: newCharacter.Description,
		Gender:      *newCharacter.Gender,
	}
	allCharacters, _ := dc.GetCharacters()
	for _, char := range allCharacters {
		if char.Name == character.Name && char.Description == character.Description {
			return nil, errors.New("character already exists")
		}
	}
	err := dc.DB.Create(&character).Error
	return &character, err
}

func (dc *DBCollection) GetCharacterByID(id int) ([]*model.Character, error) {
	var characters []*model.Character
	dc.DB.Where("id = ?", id).Find(&characters)
	return characters, nil
}

func (dc *DBCollection) GetCharacterByName(name string) ([]*model.Character, error) {
	var characters []*model.Character
	dc.DB.Where("name LIKE ?", name).Find(&characters)
	return characters, nil
}
