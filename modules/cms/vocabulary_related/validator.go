package vocabulary_related

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type VocabularyRelatedModelValidator struct {
	VocabularyID string `form:"vocabularyId" json:"vocabularyId"`
	SentenceID   string `form:"sentenceId" json:"sentenceId"`
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
