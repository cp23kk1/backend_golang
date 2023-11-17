package vocabulary_history

import (
	"cp23kk1/common/databases"
)

func CreateVocabularyHistory(userID, vocabularyID int, gameID string, correctness bool) error {
	db := databases.GetDB()
	history := &VocabularyHistoryModel{
		UserID:       userID,
		VocabularyID: vocabularyID,
		GameID:       gameID,
		Correctness:  correctness,
	}

	return db.Create(history).Error
}

func FindVocabularyHistoryByID(id int) (*VocabularyHistoryModel, error) {
	db := databases.GetDB()
	var history VocabularyHistoryModel
	if err := db.Where("id = ?", id).Preload("User").Preload("Vocabulary").First(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func FindVocabularyHistoryAll() (*[]VocabularyHistoryModel, error) {
	db := databases.GetDB()
	var history []VocabularyHistoryModel
	if err := db.Preload("User").Preload("Vocabulary").Find(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
func UpdateVocabularyHistory(id, userID int, vocabularyID int, gameID string, correctness bool) error {
	db := databases.GetDB()
	vocabularyHistory, err := FindVocabularyHistoryByID(id)
	if err != nil {
		return err
	}
	if vocabularyHistory != nil {
		history := &VocabularyHistoryModel{
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

func DeleteVocabularyHistory(history *VocabularyHistoryModel) error {
	db := databases.GetDB()
	return db.Delete(history).Error
}
