package sentence_history

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type SentenceHistoryModelValidator struct {
	UserID      int    `form:"userId" json:"userId" binding:"required"`
	SentenceID  string `form:"sentenceId" json:"sentenceId" binding:"required"`
	GameID      string `form:"gameId" json:"gameId" binding:"required,max=255"`
	AnswerID    string `form:"answerId" json:"answerId" binding:"required"`
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
