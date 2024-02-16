package users

import (
	"cp23kk1/common/databases"
	"cp23kk1/modules/repository/enum"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type UserResponse struct {
	ID             int
	Email          *string
	Role           *enum.Role
	DisplayName    *string
	Active         bool
	Image          *string
	PrivateProfile bool
	ScoreBoards    []databases.ScoreBoardModel
}

type UserSerializer struct {
	c *gin.Context
	databases.UserModel
}

func (self *UserSerializer) Response() UserResponse {
	userModel := self.UserModel
	// userModel := self.c.MustGet("my_user_model").(UserModel)
	var user UserResponse
	err := mapstructure.Decode(userModel, &user)
	if err != nil {
		fmt.Println("Error mapping data:", err)
		return user
	}
	return user
}
