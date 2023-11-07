package history

import (
	"cp23kk1/common/databases"
	"time"
)

type VocabularyHistory struct {
	ID           int    `gorm:"primary_key"`
	UserID       int    `gorm:"column:user_id"`
	VocabularyID int    `gorm:"column:vocabulary_id"`
	GameID       string `gorm:"column:game_id"`
	Correctness  bool   `gorm:"column:correctness"`
}
type SentenceHistory struct {
	ID          int    `gorm:"primary_key"`
	UserID      int    `gorm:"column:user_id"`
	SentenceID  int    `gorm:"column:sentence_id"`
	GameID      string `gorm:"column:game_id"`
	Correctness bool   `gorm:"column:correctness"`
}
type PassageHistory struct {
	ID          int    `gorm:"primary_key"`
	UserID      int    `gorm:"column:user_id"`
	PassageID   int    `gorm:"column:passage_id"`
	GameID      string `gorm:"column:game_id"`
	Correctness bool   `gorm:"column:correctness"`
}
type ScoreBoard struct {
	ID         int       `gorm:"primary_key"`
	UserID     int       `gorm:"column:user_id"`
	Score      int       `gorm:"column:score"`
	Week       int       `gorm:"column:week"`
	Start_date time.Time `gorm:"column:start_date"`
	End_date   time.Time `gorm:"column:end_date"`
}

func AutoMigrate() {
	db := databases.Init()

	db.AutoMigrate(&VocabularyHistory{})
	db.AutoMigrate(&SentenceHistory{})
	db.AutoMigrate(&PassageHistory{})
}

func (VocabularyHistory) TableName() string {
	return "vocabulary_history"
}
func (SentenceHistory) TableName() string {
	return "sentence_history"
}
func (PassageHistory) TableName() string {
	return "passage_history"
}
func (ScoreBoard) TableName() string {
	return "score_board"
}

func selectVocabularyHistoryAll() ([]VocabularyHistory, error) {
	db := databases.GetDB()
	var vocabularies []VocabularyHistory
	err := db.Raw("SELECT * FROM vocabulary_history LIMIT ?", 10).Scan(&vocabularies).Error
	return vocabularies, err
}
func selectSentenceHistoryAll() ([]SentenceHistory, error) {
	db := databases.GetDB()
	var sentences []SentenceHistory
	err := db.Raw("SELECT * FROM sentence_history LIMIT ?", 10).Scan(&sentences).Error
	return sentences, err
}
func selectPassageHistoryAll() ([]PassageHistory, error) {
	db := databases.GetDB()
	var passages []PassageHistory
	err := db.Raw("SELECT * FROM passage_history LIMIT ?", 10).Scan(&passages).Error
	return passages, err
}

// insert
func insertVocabulary(input []Vocab, gameId string) (interface{}, error) {
	db := databases.GetDB()
	var insertedId interface{}
	var err error
	for _, item := range input {
		err = db.Exec("INSERT INTO vocabulary_history (user_id, vocabulary_id, game_id, correctness) VALUES (?, ?, ?, ?)", 1, item.VocabularyID, gameId, item.Correctness).Scan(&insertedId).Error
	}
	return insertedId, err
}
