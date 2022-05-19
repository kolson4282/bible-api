package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kolson4282/bible-api/graph/generated"
	"github.com/kolson4282/bible-api/graph/model"
)

func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.NewCharacter) (*model.Character, error) {
	return r.Collection.CreateCharacter(input)
}

func (r *queryResolver) Characters(ctx context.Context, id *int, name *string) ([]*model.Character, error) {
	if id != nil {
		return r.Collection.GetCharacterByID(*id)
	}
	if name != nil {
		return r.Collection.GetCharacterByName(*name)
	}
	return r.Collection.GetCharacters()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
