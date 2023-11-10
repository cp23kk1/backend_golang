package vocabulary

type VocabularyModel struct {
	ID             int    `gorm:"primary_key"`
	Word           string `gorm:"not null;column:word;size:255"`
	Meaning        string `gorm:"not null;column:meaning;size:255"`
	Pos            string `gorm:"not null;column:pos;size:45"`
	DifficultyCefr string `gorm:"not null;column:difficulty_cefr;size:45"`
}
