package resolver_test

import (
	"errors"

	"github.com/kolson4282/bible-api/graph/model"
)

type MockCollection struct {
	characters []*model.Character
	books      []*model.Book
}

func NewMockCollection() *MockCollection {
	return &MockCollection{
		characters: []*model.Character{
			{
				ID:          1,
				Name:        "God",
				Description: "God",
			},
			{
				ID:          2,
				Name:        "Eve",
				Description: "Eve",
				Gender:      model.GenderFemale,
			},
		},
		books: []*model.Book{
			{
				ID:       1,
				Title:    "Genesis",
				Chapters: 50,
			},
		},
	}
}
func (mc *MockCollection) GetCharacters() ([]*model.Character, error) {
	return mc.characters, nil
}

func (mc *MockCollection) CreateCharacter(newCharacter model.NewCharacter) (*model.Character, error) {
	character := model.Character{
		ID:          2,
		Name:        newCharacter.Name,
		Description: newCharacter.Description,
	}
	mc.characters = append(mc.characters, &character)
	return &character, nil
}

func (mc *MockCollection) GetCharacterByID(id int) ([]*model.Character, error) {
	var characters []*model.Character
	for _, char := range mc.characters {
		if char.ID == id {
			characters = append(characters, char)
			break
		}
	}
	if len(characters) == 0 {
		return nil, errors.New("character at id not found")
	}
	return characters, nil
}

func (mc *MockCollection) GetCharacterByName(name string) ([]*model.Character, error) {
	var characters []*model.Character
	for _, char := range mc.characters {
		if char.Name == name {
			characters = append(characters, char)
			break
		}
	}
	if len(characters) == 0 {
		return nil, errors.New("character at name not found")
	}
	return characters, nil
}

func (mc *MockCollection) GetBooks() ([]*model.Book, error) {
	return mc.books, nil
}
