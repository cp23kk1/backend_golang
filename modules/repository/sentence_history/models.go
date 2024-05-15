package sentence_history

type SentenceFromGameResultModel struct {
	SentenceID  string `form:"sentenceId" json:"sentenceId" binding:"required"`
	Correctness bool   `form:"correctness" json:"correctness" binding:"required"`
	AnswerID    string `form:"answerId" json:"answerId" binding:"required"`
}

type SentenceHistoryCountModel struct {
	Count int    `json:"count"`
	Tense string `json:"tense"`
}
