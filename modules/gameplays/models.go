package gameplays

import (
	"cp23kk1/common/databases"
)

type VocabularyModel struct {
	ID             int    `gorm:"primary_key"`
	Word           string `gorm:"column:word"`
	Meaning        string `gorm:"column:meaning"`
	Pos            string `gorm:"column:pos"`
	DifficultyCefr string `gorm:"column:difficulty_cefr"`
}

func AutoMigrate() {
	db := databases.Init()

	db.AutoMigrate(&VocabularyModel{})
}

func (VocabularyModel) TableName() string {
	return "vocabulary"
}

func selectAll() ([]VocabularyModel, error) {
	db := databases.GetDB()
	var vocabularies []VocabularyModel
	err := db.Raw("SELECT * FROM vocabulary LIMIT ?", 10).Scan(&vocabularies).Error
	return vocabularies, err
}

func randomFromGamePlay() ([]VocabularyModel, error) {

	db := databases.GetDB()
	var vocabularies []VocabularyModel
	err := db.Raw("SELECT * FROM vocabulary ORDER BY RAND() LIMIT 50;").Scan(&vocabularies).Error
	return vocabularies, err
}
