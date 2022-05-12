package resolver_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/kolson4282/tdd-bible-api/graph/generated"
	"github.com/kolson4282/tdd-bible-api/graph/model"
	"github.com/kolson4282/tdd-bible-api/graph/resolver"
)

func TestCharacters(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		Collection: &MockCollection{},
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

}

type MockCollection struct {
}

func (mc *MockCollection) GetCharacters() ([]*model.Character, error) {
	return []*model.Character{
		{
			ID:          1,
			Name:        "God",
			Description: "God",
		},
	}, nil
}

func (mc *MockCollection) CreateCharacter(newCharacter model.NewCharacter) (*model.Character, error) {
	character := model.Character{
		ID:          -1,
		Name:        newCharacter.Name,
		Description: newCharacter.Description,
	}
	return &character, nil
}
