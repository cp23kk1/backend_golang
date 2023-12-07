package vocabulary

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
