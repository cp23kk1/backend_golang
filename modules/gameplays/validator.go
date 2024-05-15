package gameplays

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type MultiPlayerGetQuestionValidator struct {
	// TODO: temporarily userid future will use from request
	Mode             string `form:"mode" json:"mode"`
	NumberOfQuestion int    `form:"numberOfQuestion" json:"numberOfQuestion" binding:"required"`
}

func NewMultiPlayerValidator() MultiPlayerGetQuestionValidator {
	return MultiPlayerGetQuestionValidator{}
}

func (v *MultiPlayerGetQuestionValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
