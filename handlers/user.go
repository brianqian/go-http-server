package handlers

import (
	"base/pkg/chess_client"
	"context"
)

type UserService struct{}

type User struct {
	Id string `json:"id"`
}

func (u UserService) FindUserById(ctx context.Context, id string) User {

	chess_client.BuildChessClient()

	return User{
		Id: id,
	}
}
