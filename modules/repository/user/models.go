package user

import (
	"cp23kk1/modules/repository/enum"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID               int               `gorm:"primary_key"`
	Email            *string           `gorm:"null;column:email;size:320;"`
	Role             enum.Role         `gorm:"not null;column:role;type:enum('admin','user');"`
	DisplayName      *string           `gorm:"null;column:display_name;size:255"`
	IsActive         bool              `gorm:"not null;column:is_active"`
	Image            *string           `gorm:"null;column:image;size:255"`
	IsPrivateProfile bool              `gorm:"not null;column:is_private_profile"`
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
