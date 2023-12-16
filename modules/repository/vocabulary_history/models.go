package vocabulary_history

type VocabularyHistoryModel struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null"`
	VocabularyID uint   `gorm:"not null"`
	GameID       string `gorm:"type:varchar(45);not null"`
	Correctness  bool   `gorm:"not null"`

	// Foreign key references
	User       userModel       `gorm:"foreignKey:UserID"`
	Vocabulary VocabularyModel `gorm:"foreignKey:VocabularyID"`
}

func (VocabularyHistoryModel) TableName() string {
	return "vocabulary_history"
}

type userModel struct {
	ID               uint    `gorm:"primaryKey"`
	Email            *string `gorm:"type:varchar(320);unique;index"`
	DisplayName      *string `gorm:"type:varchar(255)"`
	IsActive         bool    `gorm:"not null;default:true"`
	Image            *string `gorm:"type:varchar(255)"`
	IsPrivateProfile bool    `gorm:"not null;default:false"`
}

func (userModel) TableName() string {
	return "user"
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

type VocabularyFromGameResultModel struct {
	VocabularyID int  `form:"vocabularyId" json:"vocabularyId" binding:"required"`
	Correctness  bool `form:"correctness" json:"correctness" binding:"required"`
}
