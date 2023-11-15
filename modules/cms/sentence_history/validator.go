package sentence_history

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type SentenceHistoryModelValidator struct {
	UserID      uint   `form:"userId" json:"userId" binding:"required"`
	SentenceID  uint   `form:"sentenceId" json:"sentenceId" binding:"required"`
	GameID      string `form:"gameId" json:"gameId" binding:"required,max=255"`
	Correctness bool   `form:"correctness" json:"correctness"`
}

func NewSentenceHistoryModelValidator() SentenceHistoryModelValidator {
	return SentenceHistoryModelValidator{}
}

func (v *SentenceHistoryModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
