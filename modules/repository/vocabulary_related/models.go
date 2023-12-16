package vocabulary_related

type VocabularyRelatedModel struct {
	VocabularyID int `gorm:"primaryKey"`
	SentenceID   int `gorm:"primaryKey"`

	Vocabulary VocabularyModel `gorm:"foreignkey:ID;references:VocabularyID"`
	Sentence   SentenceModel   `gorm:"foreignkey:ID;references:SentenceID"`
}

type VocabularyModel struct {
	ID             uint   `gorm:"primaryKey"`
	Word           string `gorm:"type:varchar(255);not null"`
	Meaning        string `gorm:"type:varchar(255);not null"`
	POS            string `gorm:"type:varchar(45);not null"`
	DifficultyCEFR string `gorm:"type:varchar(45);not null"`
}

func (VocabularyModel) TableName() string {
	return "vocabulary"
}

type SentenceModel struct {
	ID        uint   `gorm:"primaryKey"`
	PassageID *uint  `gorm:"index;foreignKey:PassageID"`
	Sequence  *int   `gorm:"index"`
	Text      string `gorm:"type:varchar(255);not null"`
	Meaning   string `gorm:"type:varchar(255);not null"`
}

func (SentenceModel) TableName() string {
	return "sentence"
}

// TableName sets the table name for the model
func (VocabularyRelatedModel) TableName() string {
	return "vocabulary_related"
}
