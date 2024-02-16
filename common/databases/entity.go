package databases

import (
	"cp23kk1/modules/repository/enum"
	"time"

	"gorm.io/gorm"
)

type PassageModel struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"type:varchar(255);not null"`
}

func (PassageModel) TableName() string {
	return "passage"
}

type PassageHistoryModel struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	PassageID   uint   `gorm:"not null"`
	GameID      string `gorm:"not null;type:varchar(45)"`
	Correctness bool   `gorm:"not null"`

	// Foreign key references
	User    UserModel    `gorm:"foreignKey:UserID"`
	Passage PassageModel `gorm:"foreignKey:PassageID"`
}

func (PassageHistoryModel) TableName() string {
	return "passage_history"
}

type ScoreBoardModel struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Score     int       `gorm:"not null"`
	Week      int       `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`

	// Foreign key reference to the User model
	User UserModel `gorm:"foreignKey:UserID"`
}

func (ScoreBoardModel) TableName() string {
	return "score_board"
}

type SentenceModel struct {
	ID        uint   `gorm:"primaryKey"`
	PassageID *uint  `gorm:"index;foreignKey:PassageID"`
	Sequence  *int   `gorm:"index"`
	Text      string `gorm:"type:varchar(255);not null"`
	Meaning   string `gorm:"type:varchar(255);not null"`

	Passage PassageModel
}

func (SentenceModel) TableName() string {
	return "sentence"
}

type SentenceHistoryModel struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	SentenceID  uint   `gorm:"not null"`
	GameID      string `gorm:"not null;type:varchar(45)"`
	Correctness bool   `gorm:"not null"`

	// Foreign key references
	User     UserModel     `gorm:"foreignKey:UserID"`
	Sentence SentenceModel `gorm:"foreignKey:SentenceID"`
}

func (SentenceHistoryModel) TableName() string {
	return "sentence_history"
}

type SentenceFromGameResultModel struct {
	SentenceID  int  `form:"sentenceId" json:"sentenceId" binding:"required"`
	Correctness bool `form:"correctness" json:"correctness" binding:"required"`
}
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

type VocabularyModel struct {
	ID             uint   `gorm:"primaryKey"`
	Word           string `gorm:"type:varchar(255);not null"`
	Meaning        string `gorm:"type:varchar(255);not null"`
	POS            string `gorm:"type:varchar(45);not null"`
	DifficultyCEFR string `gorm:"type:varchar(45);not null"`
}

func (VocabularyModel) TableName() string {
	return "vocabulary"
}

type VocabularyHistoryModel struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null"`
	VocabularyID uint   `gorm:"not null"`
	GameID       string `gorm:"type:varchar(45);not null"`
	Correctness  bool   `gorm:"not null"`

	// Foreign key references
	User       UserModel       `gorm:"foreignKey:UserID"`
	Vocabulary VocabularyModel `gorm:"foreignKey:VocabularyID"`
}

func (VocabularyHistoryModel) TableName() string {
	return "vocabulary_history"
}

type VocabularyRelatedModel struct {
	VocabularyID int `gorm:"primaryKey"`
	SentenceID   int `gorm:"primaryKey"`

	Vocabulary VocabularyModel `gorm:"foreignkey:ID;references:VocabularyID"`
	Sentence   SentenceModel   `gorm:"foreignkey:ID;references:SentenceID"`
}

// TableName sets the table name for the model
func (VocabularyRelatedModel) TableName() string {
	return "vocabulary_related"
}
