package score

import (
	"cp23kk1/common"
	"cp23kk1/modules/repository/passage_history"
	"cp23kk1/modules/repository/sentence_history"
	"cp23kk1/modules/repository/vocabulary_history"

	"github.com/gin-gonic/gin"
)

type GameResultModelValidator struct {
	// TODO: temporarily userid future will use from request
	UserID        *int                                               `form:"userID" json:"userID"`
	GameID        string                                             `form:"gameID" json:"gameID" binding:"required"`
	CurrentSocore int                                                `form:"current_score" json:"current_score" binding:"required"`
	Vocabs        []vocabulary_history.VocabularyFromGameResultModel `form:"vocabs" json:"vocabs" binding:"required"`
	Sentences     []sentence_history.SentenceFromGameResultModel     `form:"sentences" json:"sentences" binding:"required"`
	Passages      []passage_history.PassageFromGameResultModel       `form:"passages" json:"passages" binding:"required"`
}

func NewGameResultModelValidator() GameResultModelValidator {
	return GameResultModelValidator{}
}

func (v *GameResultModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
