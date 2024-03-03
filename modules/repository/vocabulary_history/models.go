package vocabulary_history

type VocabularyFromGameResultModel struct {
	VocabularyID int  `form:"vocabularyId" json:"vocabularyId" binding:"required"`
	Correctness  bool `form:"correctness" json:"correctness" binding:"required"`
}
