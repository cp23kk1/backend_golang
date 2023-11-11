package vocabulary

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type VocabularyModelValidator struct {
	Word           string `form:"word" json:"word" binding:"required,max=255"`
	Meaning        string `form:"meaning" json:"meaning" binding:"required"`
	Pos            string `form:"pos" json:"pos" binding:"required,max=255"`
	DifficultyCefr string `form:"difficultyCefr" json:"difficultyCefr" binding:"required,max=255"`
}

func NewVocabularyModelValidator() VocabularyModelValidator {
	return VocabularyModelValidator{}
}

func (v *VocabularyModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
