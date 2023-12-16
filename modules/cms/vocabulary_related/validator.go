package vocabulary_related

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type VocabularyRelatedModelValidator struct {
	VocabularyID int `form:"vocabularyId" json:"vocabularyId"`
	SentenceID   int `form:"sentenceId" json:"sentenceId"`
}

func NewVocabularyRelatedModelValidator() VocabularyRelatedModelValidator {
	return VocabularyRelatedModelValidator{}
}

func (v *VocabularyRelatedModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
