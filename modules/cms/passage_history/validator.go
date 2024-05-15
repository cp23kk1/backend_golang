package passage_history

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type PassageHistoryModelValidator struct {
	UserID      int    `form:"userId" json:"userId" binding:"required"`
	PassageID   string `form:"PassageId" json:"PassageId" binding:"required"`
	GameID      string `form:"gameId" json:"gameId" binding:"required,max=255"`
	Correctness bool   `form:"correctness" json:"correctness"`
}

func NewPassageHistoryModelValidator() PassageHistoryModelValidator {
	return PassageHistoryModelValidator{}
}

func (v *PassageHistoryModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
