package multiplayer

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type UpdateLobbyModelValidator struct {
	// TODO: temporarily userid future will use from request
	RoomID   string `form:"roomID" json:"roomID" binding:"required"`
	IsPlayed bool   `form:"isPlayed" json:"isPlayed" `
}

func NewUpdateLobbyModelValidator() UpdateLobbyModelValidator {
	return UpdateLobbyModelValidator{}
}

func (v *UpdateLobbyModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
