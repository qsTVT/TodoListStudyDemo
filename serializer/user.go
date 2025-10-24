package serializer

import (
	"golang/TodoList/model"
)

type User struct {
	ID        uint   `json:"id" form:"id" example:"1"`
	UserName  string `json:"user_name" form:"user_name" example:"FanOne"`
	Status    string `json:"status" form:"status"`
	CreatedAt int64  `json:"created_at" form:"created_at"`
}

func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
