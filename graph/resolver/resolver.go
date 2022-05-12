package resolver

import "github.com/kolson4282/tdd-bible-api/graph/model"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Collection interface {
	GetCharacters() ([]*model.Character, error)
	GetCharacterByID(id int) ([]*model.Character, error)
	GetCharacterByName(name string) ([]*model.Character, error)
	CreateCharacter(model.NewCharacter) (*model.Character, error)
}

type Resolver struct {
	Collection Collection
}
