package resolv

import (
	"context"

	"github.com/vkevv/go-graphql/src/graph/model"
)

// Register to store a new user
func (r *Res) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	_, err := r.UserTx.GetUserByEmail(input.Email)
	if err == nil {
		return nil, err
	}

	user := &model.User{
		Name:  input.Name,
		Email: input.Email,
	}
	err = user.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	err = r.UserTx.Create(user)
	if err != nil {
		return nil, err
	}
	token, err := r.GenToken(user.ID)
	if err != nil {
		return nil, err
	}
	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

// Login handles login
func (r *Res) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := r.UserTx.GetUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	err = user.CheckPassword(input.Password)
	if err != nil {
		return nil, err
	}
	token, err := r.GenToken(user.ID)
	if err != nil {
		return nil, err
	}
	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
