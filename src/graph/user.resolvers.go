package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vkevv/go-graphql/src/graph/generated"
	"github.com/vkevv/go-graphql/src/graph/model"
)

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
