package sentence

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type SentenceModelValidator struct {
	PassageID *string `form:"passageId" json:"passageId"`
	Sequence  *int    `form:"sequence" json:"sequence"`
	Text      string  `form:"text" json:"text" binding:"required"`
	Meaning   string  `form:"meaning" json:"meaning" binding:"required"`
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
