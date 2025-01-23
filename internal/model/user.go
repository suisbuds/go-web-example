package model

import (
	"github.com/suisbuds/miao/pkg/app"
)

type User struct {
	*Model
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	UserType uint8  `json:"user_type"`
}

func (u User) TableName() string {
	return "mio_user"
}

type UserSwagger struct {
	List  []*User
	Pager *app.Pager
}
