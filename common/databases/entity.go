package databases

import (
	"cp23kk1/common/enum"
	"time"

	"gorm.io/gorm"
)

type PassageModel struct {
	ID           string `gorm:"primaryKey"`
	Title        string `gorm:"type:varchar(256);not null"`
	DifficultyID *uint  `gorm:"null"`

	// Foreign key reference to the Difficulty model
	Difficulty DifficultyModel `gorm:"foreignKey:DifficultyID"`
	Sentences  []SentenceModel `gorm:"foreignKey:PassageID;references:id"`
}

func (PassageModel) TableName() string {
	return "passage"
}

type PassageHistoryModel struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null"`
	PassageID    string `gorm:"not null"`
	SentenceID   string `gorm:"not null"`
	VocabularyID string `gorm:"not null"`
	GameID       string `gorm:"not null;type:varchar(45)"`
	Correctness  bool   `gorm:"not null"`

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
	GameID    string    `gorm:"type:varchar(36);not null"`
	Mode      string    `gorm:"type:varchar(256);not null"`

	// Foreign key reference to the User model
	User UserModel `gorm:"foreignKey:UserID"`
}

func (ScoreBoardModel) TableName() string {
	return "score_board"
}

type DifficultyModel struct {
	ID          uint   `gorm:"primaryKey"`
	Standard    string `gorm:"type:varchar(16);not null"`
	Level       string `gorm:"type:varchar(16);not null"`
	Description string `gorm:"type:varchar(256);not null"`
}

func (DifficultyModel) TableName() string {
	return "difficulty"
}

type SentenceModel struct {
	ID        string  `gorm:"primaryKey"`
	PassageID *string `gorm:"index;foreignKey:PassageID"`
	Sequence  *int    `gorm:"index"`
	Sentence  string  `gorm:"type:varchar(512);not null"`
	Meaning   string  `gorm:"type:varchar(1024);not null"`
	Tense     string  `gorm:"type:varchar(512);not null"`

	Vocabularies []*VocabularyModel `gorm:"many2many:vocabulary_related;foreignKey:ID;joinForeignKey:sentence_id;References:ID;JoinReferences:vocabulary_id;"`
	Passage      PassageModel
}

func (SentenceModel) TableName() string {
	return "sentence"
}

type SentenceHistoryModel struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null"`
	SentenceID   string `gorm:"not null"`
	GameID       string `gorm:"not null;type:varchar(45)"`
	Correctness  bool   `gorm:"not null"`
	VocabularyID string `gorm:"not null"`

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

	ID               uint      `gorm:"primaryKey"`
	Email            *string   `gorm:"type:varchar(320);unique;index"`
	Role             enum.Role `gorm:"not null;column:role;type:enum('admin','user','guest');"`
	DisplayName      *string   `gorm:"type:varchar(255)"`
	IsActive         bool      `gorm:"not null;default:true"`
	Image            *string   `gorm:"type:varchar(255)"`
	IsPrivateProfile bool      `gorm:"not null;default:false"`
	CreatedAt        time.Time
	ScoreBoards      []ScoreBoardModel `gorm:"foreignKey:UserID;references:id"`
}

func (UserModel) TableName() string {
	return "user"
}

type VocabularyModel struct {
	ID           string `gorm:"primaryKey"`
	DifficultyID uint   `gorm:"null"`
	Vocabulary   string `gorm:"type:varchar(256);not null"`
	Meaning      string `gorm:"type:varchar(256);not null"`
	Definition   string `gorm:"type:varchar(256);not null"`
	POS          string `gorm:"type:varchar(256);not null"`
	Tag          string `gorm:"type:varchar(256);not null"`
	Lemma        string `gorm:"type:varchar(256);not null"`
	Dep          string `gorm:"type:varchar(256);not null"`

	Sentences []*SentenceModel `gorm:"many2many:vocabulary_related;foreignKey:ID;joinForeignKey:vocabulary_id;References:ID;JoinReferences:sentence_id;"`

	// Foreign key reference to the Difficulty model
	Difficulty DifficultyModel `gorm:"foreignKey:DifficultyID"`
}

func (VocabularyModel) TableName() string {
	return "vocabulary"
}

type VocabularyHistoryModel struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null"`
	VocabularyID string `gorm:"not null"`
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
	VocabularyID string `gorm:"primaryKey"`
	SentenceID   string `gorm:"primaryKey"`

	Vocabulary VocabularyModel `gorm:"foreignkey:ID;references:VocabularyID"`
	Sentence   SentenceModel   `gorm:"foreignkey:ID;references:SentenceID"`
}

// TableName sets the table name for the model
func (VocabularyRelatedModel) TableName() string {
	return "vocabulary_related"
}
