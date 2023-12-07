package user

import (
	"cp23kk1/modules/repository/enum"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model

	ID               uint              `gorm:"primaryKey"`
	Email            *string           `gorm:"type:varchar(320);unique;index"`
	Role             enum.Role         `gorm:"not null;column:role;type:enum('admin','user');"`
	DisplayName      *string           `gorm:"type:varchar(255)"`
	IsActive         bool              `gorm:"not null;default:true"`
	Image            *string           `gorm:"type:varchar(255)"`
	IsPrivateProfile bool              `gorm:"not null;default:false"`
	ScoreBoards      []ScoreBoardModel `gorm:"foreignKey:UserID;references:id"`
}

func (UserModel) TableName() string {
	return "user"
}

type ScoreBoardModel struct {
	ID        int
	UserID    int `gorm:"not null"`
	Score     int
	Week      int
	StartDate time.Time
	EndDate   time.Time
}

func (ScoreBoardModel) TableName() string {
	return "score_board"
}
