package score_board

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type ScoreBoardModelValidator struct {
	UserID    int    `form:"userId" json:"userId" binding:"required"`
	Score     int    `form:"score" json:"score" binding:"required"`
	Week      int    `form:"week" json:"week" binding:"required"`
	StartDate string `form:"startDate" json:"startDate" binding:"required" `
	EndDate   string `form:"endDate" json:"endDate" binding:"required" `
	GameID    string `form:"gameId" json:"gameId" binding:"required" `
	Mode      string `form:"mode" json:"mode" binding:"required" `
}

func NewScoreBoardModelValidator() ScoreBoardModelValidator {
	return ScoreBoardModelValidator{}
}

func (v *ScoreBoardModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
