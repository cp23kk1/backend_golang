package vocabulary_history

type VocabularyFromGameResultModel struct {
	VocabularyID string `form:"vocabularyId" json:"vocabularyId" binding:"required"`
	Correctness  bool   `form:"correctness" json:"correctness" binding:"required"`
}
