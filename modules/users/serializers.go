package users

import (
	"cp23kk1/common/databases"
	"cp23kk1/common/enum"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	ID               uint      `json:"id"`
	Email            *string   `json:"email"`
	Role             enum.Role `json:"role"`
	DisplayName      *string   `json:"displayName"`
	IsActive         bool      `json:"isActive"`
	Image            *string   `json:"image"`
	CreateAt         string    `json:"createAt"`
	IsPrivateProfile bool      `json:"isPrivate"`
}

type UserSerializer struct {
	c *gin.Context
	databases.UserModel
}

func ConvertToUserResponse(user databases.UserModel) UserResponse {

	return UserResponse{
		ID:               user.ID,
		Email:            user.Email,
		Role:             user.Role,
		DisplayName:      user.DisplayName,
		IsActive:         user.IsActive,
		Image:            user.Image,
		CreateAt:         user.CreatedAt.String(),
		IsPrivateProfile: user.IsPrivateProfile,
	}
}
func (self *UserSerializer) Response() UserResponse {
	userModel := self.UserModel
	// userModel := self.c.MustGet("my_user_model").(UserModel)

	return ConvertToUserResponse(userModel)
}
