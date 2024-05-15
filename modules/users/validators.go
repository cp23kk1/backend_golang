package users

import (
	"cp23kk1/common"
	"cp23kk1/common/databases"

	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"exists,email"`
	} `json:"user"`
	userModel databases.UserModel `json:"-"`
}

type ChangeDisplayNameValidator struct {
	NewDisplayName string ` json:"displayName" binding:"required,max=255"`
}

func NewChangeDisplayNameValidator() ChangeDisplayNameValidator {
	changeDisplayNameValidator := ChangeDisplayNameValidator{}
	//userModelValidator.User.Email ="w@g.cn"
	return changeDisplayNameValidator
}

func (v *ChangeDisplayNameValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}

// func (self *UserModelValidator) Bind(c *gin.Context) error {
// 	err := common.Bind(c, self)
// 	if err != nil {
// 		return err
// 	}
// 	self.userModel.Username = self.User.Username
// 	self.userModel.Email = self.User.Email

// 	return nil
// }

func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	//userModelValidator.User.Email ="w@g.cn"
	return userModelValidator
}
