package resolver_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/kolson4282/tdd-bible-api/graph/generated"
	"github.com/kolson4282/tdd-bible-api/graph/resolver"
)

func TestCharacters(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))

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

}
