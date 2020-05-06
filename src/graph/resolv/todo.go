package resolv

import (
	"context"
	"errors"

	"github.com/vkevv/go-graphql/src/graph/model"
	"github.com/vkevv/go-graphql/src/middleware"
)

// GetUserTodos userTodos
func (r *Res) GetUserTodos(ctx context.Context) ([]*model.Todo, error) {
	userID, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return r.TodoTx.FromUserID(userID)
}

// CreateTodo create a todo
func (r *Res) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	errorShort := errors.New("Argument too short")
	if len(input.Text) < 5 {
		return nil, errorShort
	}
	userID, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	todo := model.Todo{
		Text:   input.Text,
		Done:   false,
		UserID: userID,
	}
	err = r.TodoTx.Create(&todo)
	return &todo, err
}
