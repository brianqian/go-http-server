package services

import (
	"context"
)

type UserService struct{}

type User struct {
	Id string `json:"id"`
}

func (u UserService) FindUserById(ctx context.Context, id string) User {

	return User{
		Id: id,
	}
}
