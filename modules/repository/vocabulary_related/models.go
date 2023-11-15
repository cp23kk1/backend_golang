package vocabulary_related

type VocabularyRelatedModel struct {
	VocabularyID int `gorm:"primaryKey"`
	SentenceID   int `gorm:"primaryKey"`

	// Add other fields if necessary

	// Relationships
	Vocabulary VocabularyModel `gorm:"foreignkey:ID;references:VocabularyID"`
	Sentence   SentenceModel   `gorm:"foreignkey:ID;references:SentenceID"`
}

type VocabularyModel struct {
	ID             int
	Word           string
	Meaning        string
	Pos            string
	DifficultyCefr string
}

func (VocabularyModel) TableName() string {
	return "vocabulary"
}

type SentenceModel struct {
	ID        int
	PassageID int
	Sequence  int
	Text      string
	Meaning   string
}

func (SentenceModel) TableName() string {
	return "sentence"
}

// TableName sets the table name for the model
func (VocabularyRelatedModel) TableName() string {
	return "vocabulary_related"
}
