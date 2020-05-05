package db

import (
	"github.com/go-pg/pg/v9"
	"github.com/vkevv/go-graphql/src/graph/model"
)

// TodoTx handles all Todo database transactions
type TodoTx struct {
	DB *pg.DB
}

// Create stores a todo in DB
func (t *TodoTx) Create(todo *model.Todo) error {
	_, err := t.DB.Model(todo).Returning("*").Insert()
	return err
}

// FromUserID get Todos from userid
func (t *TodoTx) FromUserID(userID string) ([]*model.Todo, error) {
	todos := make([]*model.Todo, 0)
	err := t.DB.Model(&todos).Where("user_id = ?", userID).Order("id").Select()
	return todos, err
}

// GetAll get all todos
func (t *TodoTx) GetAll() ([]*model.Todo, error) {
	todos := make([]*model.Todo, 0)
	err := t.DB.Model(&todos).Order("id").Select()
	return todos, err
}
