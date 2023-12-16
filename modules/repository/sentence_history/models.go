package sentence_history

type SentenceHistoryModel struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	SentenceID  uint   `gorm:"not null"`
	GameID      string `gorm:"not null;type:varchar(45)"`
	Correctness bool   `gorm:"not null"`

	// Foreign key references
	User     userModel     `gorm:"foreignKey:UserID"`
	Sentence SentenceModel `gorm:"foreignKey:SentenceID"`
}

func (SentenceHistoryModel) TableName() string {
	return "sentence_history"
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

type SentenceFromGameResultModel struct {
	SentenceID  int  `form:"sentenceId" json:"sentenceId" binding:"required"`
	Correctness bool `form:"correctness" json:"correctness" binding:"required"`
}

type SentenceModel struct {
	ID        int
	PassageID int
	Sequence  int
	Text      string
	Meaning   string
}

func (SentenceModel) TableName() string {
	return "sentence"
}
