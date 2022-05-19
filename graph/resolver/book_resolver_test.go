package resolver_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/kolson4282/bible-api/graph/generated"
	"github.com/kolson4282/bible-api/graph/resolver"
)

func TestBook(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		Collection: NewMockCollection(),
	}})))

	t.Run("Get All Books", func(t *testing.T) {
		var resp struct {
			Books []struct {
				ID       int
				Title    string
				Chapters int
			}
		}
		q := `
		{
			books {
				id
				title
				chapters
			}
		}`
		c.MustPost(q, &resp)
		if len(resp.Books) == 0 {
			t.Error("no Books returned")
		}
		for i, book := range resp.Books {
			if book.Title == "" {
				t.Errorf("Invalid book title at postion %d", i)
			}
		}
	})
}
