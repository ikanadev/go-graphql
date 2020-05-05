package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vkevv/go-graphql/src/graph/generated"
	"github.com/vkevv/go-graphql/src/graph/model"
)

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return r.Res.UserTx.GetUserByID(obj.UserID)
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
