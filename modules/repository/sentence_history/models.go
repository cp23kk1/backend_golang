package sentence_history

type SentenceFromGameResultModel struct {
	SentenceID  int  `form:"sentenceId" json:"sentenceId" binding:"required"`
	Correctness bool `form:"correctness" json:"correctness" binding:"required"`
}
