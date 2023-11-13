package vocabulary_history

import (
	"cp23kk1/common/databases"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {

	db.AutoMigrate(&VocabularyHistory{})
}

func CreateVocabularyHistory(userID uint, vocabularyID uint, gameID string, correctness bool) error {
	db := databases.GetDB()
	history := &VocabularyHistory{
		UserID:       userID,
		VocabularyID: vocabularyID,
		GameID:       gameID,
		Correctness:  correctness,
	}

	return db.Create(history).Error
}

func GetVocabularyHistoryByID(id uint) (*VocabularyHistory, error) {
	db := databases.GetDB()
	var history VocabularyHistory
	if err := db.Where("id = ?", id).Preload("User").Preload("Vocabulary").First(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func GetVocabularyHistoryAll() (*[]VocabularyHistory, error) {
	db := databases.GetDB()
	var history []VocabularyHistory
	if err := db.Preload("User").Preload("Vocabulary").Find(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
func UpdateVocabularyHistory(id uint, userID uint, vocabularyID uint, gameID string, correctness bool) error {
	db := databases.GetDB()
	vocabularyHistory, err := GetVocabularyHistoryByID(id)
	if err != nil {
		return err
	}
	if vocabularyHistory != nil {
		history := &VocabularyHistory{
			ID:           id,
			UserID:       userID,
			VocabularyID: vocabularyID,
			GameID:       gameID,
			Correctness:  correctness,
		}
		return db.Save(history).Error
	}
	return nil
}

func DeleteVocabularyHistory(history *VocabularyHistory) error {
	db := databases.GetDB()
	return db.Delete(history).Error
}
