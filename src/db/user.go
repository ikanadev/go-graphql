package db

import (
	"github.com/go-pg/pg/v9"
	"github.com/vkevv/go-graphql/src/graph/model"
)

// UserTx handles all User database transactions.
type UserTx struct {
	DB *pg.DB
}

// GetUserByID get user by id
func (u *UserTx) GetUserByID(id string) (*model.User, error) {
	user := model.User{}
	err := u.DB.Model(&user).Where("id = ?", id).First()
	return &user, err
}

// GetUserByEmail Get user by email
func (u *UserTx) GetUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := u.DB.Model(&user).Where("email = ?", email).First()
	return &user, err
}

// Create stores an user
func (u *UserTx) Create(user *model.User) error {
	_, err := u.DB.Model(user).Returning("*").Insert()
	return err
}
