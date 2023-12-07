package score_board

import (
	"cp23kk1/modules/repository/enum"
	"time"
)

type ScoreBoardModel struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Score     int       `gorm:"not null"`
	Week      int       `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`

	// Foreign key reference to the User model
	User userModel `gorm:"foreignKey:UserID"`
}

func (ScoreBoardModel) TableName() string {
	return "score_board"
}

type userModel struct {
	ID               uint      `gorm:"primaryKey"`
	Email            *string   `gorm:"type:varchar(320);unique;index"`
	Role             enum.Role `gorm:"not null;column:role;type:enum('admin','user');"`
	DisplayName      *string   `gorm:"type:varchar(255)"`
	IsActive         bool      `gorm:"not null;default:true"`
	Image            *string   `gorm:"type:varchar(255)"`
	IsPrivateProfile bool      `gorm:"not null;default:false"`
}

func (userModel) TableName() string {
	return "user"
}
