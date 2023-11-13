package score_board

import (
	"cp23kk1/modules/repository/enum"
	"time"
)

type ScoreBoardModel struct {
	ID        int `gorm:"primary_key"`
	UserID    int `gorm:"not null"`
	User      userModel
	Score     int       `gorm:"not null"`
	Week      int       `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
}
type userModel struct {
	ID               int
	Email            *string `gorm:"null;column:email;size:255;"`
	Role             enum.Role
	DisplayName      *string `gorm:"null;column:display_name;size:255;"`
	IsActive         bool
	Image            *string
	IsPrivateProfile bool
}

func (userModel) TableName() string {
	return "user"
}
