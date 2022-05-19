package resolver_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/kolson4282/bible-api/graph/generated"
	"github.com/kolson4282/bible-api/graph/resolver"
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
				Gender      string
			}
		}
		q := `
		{
			characters {
				id
				name
				description
				gender
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
			createCharacter(input: {name: "Adam", description: "First Man"}) {
			  id
			  name
			  description
			}
		  }`
		c.MustPost(q, &resp)
		if resp.CreateCharacter.Name != "Adam" {
			t.Errorf("Name added incorrectly, got %s, want %s", resp.CreateCharacter.Name, "Adam")
		}
		if resp.CreateCharacter.Description != "First Man" {
			t.Errorf("Description added incorrectly, got %s, want %s", resp.CreateCharacter.Description, "First Man")
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

	t.Run("Get Characters by Name", func(t *testing.T) {
		var resp struct {
			Characters []struct {
				ID          int
				Name        string
				Description string
			}
		}
		q := `
		{
			characters (name: "Adam"){
				id
				name
				description
			}
		}`
		c.MustPost(q, &resp)
		if len(resp.Characters) == 0 {
			t.Fatalf("No Characters Returned")
		}
		if resp.Characters[0].Name != "Adam" {
			t.Errorf("got the wrong Name. Expected Adam, got %s", resp.Characters[0].Name)
		}
	})
}
