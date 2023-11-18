package sentence

type SentenceModel struct {
	ID        int `gorm:"primaryKey"`
	PassageID *int
	Sequence  *int
	Text      string
	Meaning   string

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
