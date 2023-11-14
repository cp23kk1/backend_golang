package sentence

type SentenceModel struct {
	ID        int    `gorm:"primaryKey"`
	PassageID int    `gorm:"index"`
	Sequence  int    `gorm:"not null"`
	Text      string `gorm:"not null"`
	Meaning   string `gorm:"not null"`

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
