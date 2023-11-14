package user

import (
	"cp23kk1/common"
	"cp23kk1/modules/repository/enum"

	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	Email            string    `form:"email" json:"email"`
	Role             enum.Role `form:"role" json:"role" binding:"required"`
	DisplayName      string    `form:"displayName" json:"displayName" binding:"required"`
	Image            string    `form:"image" json:"image" binding:"required"`
	IsPrivateProfile bool      `form:"isPrivateProfile" json:"isPrivateProfile" `
}

func NewUserModelValidator() UserModelValidator {
	return UserModelValidator{}
}

func (v *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
