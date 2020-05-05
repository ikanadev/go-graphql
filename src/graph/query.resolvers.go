package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vkevv/go-graphql/src/graph/generated"
	"github.com/vkevv/go-graphql/src/graph/model"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.Res.GetUserTodos(ctx)
}

func (r *queryResolver) AllTodos(ctx context.Context) ([]*model.Todo, error) {
	return r.Res.TodoTx.GetAll()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
