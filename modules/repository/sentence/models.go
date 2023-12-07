package sentence

type SentenceModel struct {
	ID        uint   `gorm:"primaryKey"`
	PassageID *uint  `gorm:"index;foreignKey:PassageID"`
	Sequence  *int   `gorm:"index"`
	Text      string `gorm:"type:varchar(255);not null"`
	Meaning   string `gorm:"type:varchar(255);not null"`

	Passage PassageModel
}

type PassageModel struct {
	ID    int
	Title string
}

func (PassageModel) TableName() string {
	return "passage"
}

func (SentenceModel) TableName() string {
	return "sentence"
}
