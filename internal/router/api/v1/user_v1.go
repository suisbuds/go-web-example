package v1

import "github.com/gin-gonic/gin"


type User struct{}

func NewUser() User {
	return User{}
}

func (u *User) CreateUser(c *gin.Context) {

}

func (u *User) GetUser(c *gin.Context) {

}

func (u *User) GetUserList(c *gin.Context) {

}

func (u *User) UpdateUser(c *gin.Context) {}


func (u *User) DeleteUser(c *gin.Context) {

}

func (u *User) Logout(c *gin.Context) {

}