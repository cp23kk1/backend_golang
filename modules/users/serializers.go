package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)


type UserSerializer struct {
	c *gin.Context
	UserModel
}

type UserResponse struct {
	Username	string
	Email       string
}

func (self *UserSerializer) Response() (UserResponse, error) {
	userModel := self.UserModel
	// userModel := self.c.MustGet("my_user_model").(UserModel)
	var user UserResponse
	err := mapstructure.Decode(userModel, &user)
	if err != nil {
		fmt.Println("Error mapping data:", err)
		return user, err
	}
	return user, nil
}
