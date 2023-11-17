package vocabulary_history

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type VocabularyHistoryModelValidator struct {
	UserID       int    `form:"userId" json:"userId" binding:"required"`
	VocabularyID int    `form:"vocabularyId" json:"vocabularyId" binding:"required"`
	GameID       string `form:"gameId" json:"gameId" binding:"required,max=255"`
	Correctness  bool   `form:"correctness" json:"correctness"`
}

func NewVocabularyHistoryModelValidator() VocabularyHistoryModelValidator {
	return VocabularyHistoryModelValidator{}
}

func (v *VocabularyHistoryModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
