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
	var resp struct {
		Characters []struct {
			ID int
		}
	}
	q := `{
		characters {
			id
		}
		}`
	c.MustPost(q, &resp)
	if resp.Characters[0].ID != 0 {
		t.Error("error")
	}
}
