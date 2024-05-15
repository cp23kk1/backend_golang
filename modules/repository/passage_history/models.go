package passage_history

type PassageFromGameResultModel struct {
	PassageID    string `form:"passageId" json:"passageId" binding:"required"`
	SentenceID   string `form:"sentenceId" json:"sentenceId" binding:"required"`
	VocabularyID string `form:"vocabularyId" json:"vocabularyId" binding:"required"`
	Correctness  bool   `form:"correctness" json:"correctness" binding:"required"`
}
