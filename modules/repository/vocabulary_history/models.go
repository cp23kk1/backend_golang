package vocabulary_history

type VocabularyHistory struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	VocabularyID uint
	GameID       string
	Correctness  int8

	User       userModel       `gorm:"foreignKey:UserID"`
	Vocabulary VocabularyModel `gorm:"foreignKey:VocabularyID"`
}

func (VocabularyHistory) TableName() string {
	return "vocabulary_history"
}

type userModel struct {
	ID             int
	Email          *string
	DisplayName    *string
	Active         bool
	Image          *string
	PrivateProfile bool
}

func (userModel) TableName() string {
	return "user"
}

type VocabularyModel struct {
	ID             int
	Word           string
	Meaning        *string
	Pos            string
	DifficultyCefr string
}

func (VocabularyModel) TableName() string {
	return "vocabulary"
}
