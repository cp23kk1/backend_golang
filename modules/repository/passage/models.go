package passage

type PassageModel struct {
	ID    int    `gorm:"primary_key"`
	Title string `gorm:"not null;column:title;size:255;"`
}
