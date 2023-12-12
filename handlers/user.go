package handlers

import (
	"base/types"
	"context"
	"fmt"
)

type UserService struct{}

type User struct {
	Id string `json:"id"`
}

func (u UserService) FindUserById(ctx context.Context, id string) User {

	last2 := ctx.Value(types.UserIdKey)

	fmt.Println("VAL", last2)
	return User{
		Id: id,
	}
}
