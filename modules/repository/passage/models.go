package passage

type PassageModel struct {
	ID    int `gorm:"primary_key"`
	Title string
}

func (PassageModel) TableName() string {
	return "passage"
}
