package passage_history

type PassageHistoryModel struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	PassageID   uint   `gorm:"not null"`
	GameID      string `gorm:"not null;type:varchar(45)"`
	Correctness bool   `gorm:"not null"`

	// Foreign key references
	User    userModel    `gorm:"foreignKey:UserID"`
	Passage PassageModel `gorm:"foreignKey:PassageID"`
}

func (PassageHistoryModel) TableName() string {
	return "passage_history"
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

type PassageModel struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"type:varchar(255);not null"`
}

func (PassageModel) TableName() string {
	return "passage"
}

type PassageFromGameResultModel struct {
	PassageID   int  `form:"passageId" json:"passageId" binding:"required"`
	Correctness bool `form:"correctness" json:"correctness" binding:"required"`
}
