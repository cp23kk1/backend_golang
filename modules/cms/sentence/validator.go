package sentence

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type SentenceModelValidator struct {
	PassageID int    `form:"passageId" json:"passageId" binding:"required"`
	Sequence  int    `form:"sequence" json:"sequence" binding:"required"`
	Text      string `form:"text" json:"text" binding:"required"`
	Meaning   string `form:"meaning" json:"meaning" binding:"required"`
}

func NewSentenceModelValidator() SentenceModelValidator {
	return SentenceModelValidator{}
}

func (v *SentenceModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
