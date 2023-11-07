package history

type GameResultRequest struct {
	GameID       string     `json:"gameId"`
	CurrentScore string     `json:"current_score"`
	Vocabs       []Vocab    `json:"vocabs"`
	Sentences    []Sentence `json:"sentences"`
	Passages     []Passage  `json:"passages"`
}

type Vocab struct {
	ID           string `json:"id"`
	VocabularyID string `json:"vocabularyId"`
	Correctness  bool   `json:"correctness"`
}

type Sentence struct {
	ID          string `json:"id"`
	SentenceID  string `json:"sentenceId"`
	Correctness bool   `json:"correctness"`
}

type Passage struct {
	ID          string `json:"id"`
	PassageID   string `json:"passageId"`
	Correctness bool   `json:"correctness"`
}
