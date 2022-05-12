package resolver_test

import (
	"errors"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/kolson4282/tdd-bible-api/graph/generated"
	"github.com/kolson4282/tdd-bible-api/graph/model"
	"github.com/kolson4282/tdd-bible-api/graph/resolver"
)

func TestCharacters(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		Collection: NewMockCollection(),
	}})))

	t.Run("Get All Characters", func(t *testing.T) {
		var resp struct {
			Characters []struct {
				ID          int
				Name        string
				Description string
			}
		}
		q := `
		{
			characters {
				id
				name
				description
			}
		}`
		c.MustPost(q, &resp)
		if len(resp.Characters) == 0 {
			t.Error("no characters returned")
		}
		for i, char := range resp.Characters {
			if char.Name == "" {
				t.Errorf("Invalid character name at postion %d", i)
			}
		}
	})

	t.Run("Create Character", func(t *testing.T) {
		var resp struct {
			CreateCharacter struct {
				ID          int
				Name        string
				Description string
			}
		}
		q := `
		mutation{
			createCharacter(input: {name: "adam", description: "first man"}) {
			  id
			  name
			  description
			}
		  }`
		c.MustPost(q, &resp)
		if resp.CreateCharacter.Name != "adam" {
			t.Errorf("Name added incorrectly, got %s, want %s", resp.CreateCharacter.Name, "adam")
		}
		if resp.CreateCharacter.Description != "first man" {
			t.Errorf("Description added incorrectly, got %s, want %s", resp.CreateCharacter.Description, "first man")
		}
	})

	t.Run("Get Characters by ID", func(t *testing.T) {
		var resp struct {
			Characters []struct {
				ID          int
				Name        string
				Description string
			}
		}
		q := `
		{
			characters (id: 1){
				id
				name
				description
			}
		}`
		c.MustPost(q, &resp)
		if len(resp.Characters) != 1 {
			t.Errorf("incorrect number of characters returned. Expected 1, got %d", len(resp.Characters))
		}
		if resp.Characters[0].ID != 1 {
			t.Errorf("got the wrong ID. Expected 1, got %d", resp.Characters[0].ID)
		}
	})
}

type MockCollection struct {
	characters []*model.Character
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
