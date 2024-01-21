package api

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

	userCtx := ctx.Value(types.UserIdKey)
	fmt.Println("user context: ", userCtx)
	return User{
		Id: id,
	}
}
