package passage

type PassageModel struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"type:varchar(255);not null"`
}

func (PassageModel) TableName() string {
	return "passage"
}
