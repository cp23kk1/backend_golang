package score_board

import (
	"cp23kk1/modules/repository/enum"
	"time"
)

type ScoreBoardModel struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	User      userModel
	Score     int       `gorm:"not null"`
	Week      int       `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
}
type userModel struct {
	ID               int `gorm:"primary_key"`
	Email            *string
	Role             enum.Role
	DisplayName      *string
	IsActive         bool
	Image            *string
	IsPrivateProfile bool
}

func (userModel) TableName() string {
	return "user"
}
