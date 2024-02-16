package passage_history

type PassageFromGameResultModel struct {
	PassageID   int  `form:"passageId" json:"passageId" binding:"required"`
	Correctness bool `form:"correctness" json:"correctness" binding:"required"`
}
