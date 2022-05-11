package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kolson4282/tdd-bible-api/graph/generated"
	"github.com/kolson4282/tdd-bible-api/graph/model"
)

func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.NewCharacter) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	return r.Collection.GetCharacters()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
