package vocabulary_history

type VocabularyHistoryModel struct {
	ID           int `gorm:"primaryKey"`
	UserID       int
	VocabularyID int
	GameID       string
	Correctness  bool

	User       userModel
	Vocabulary VocabularyModel
}

func (VocabularyHistoryModel) TableName() string {
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
