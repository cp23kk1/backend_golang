package sentence_history

import (
	"cp23kk1/common/databases"
)

// func AutoMigrate(db *gorm.DB) {

// 	db.AutoMigrate(&SentenceHistoryModel{})
// }

func CreateSentenceHistory(userID uint, sentenceID uint, gameID string, correctness bool) error {
	db := databases.GetDB()
	history := &SentenceHistoryModel{
		UserID:      userID,
		SentenceID:  sentenceID,
		GameID:      gameID,
		Correctness: correctness,
	}

	return db.Create(history).Error
}

func GetSentenceHistoryByID(id uint) (*SentenceHistoryModel, error) {
	db := databases.GetDB()
	var history SentenceHistoryModel
	if err := db.Where("id = ?", id).Preload("User").Preload("Sentence").First(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func GetSentenceHistoryAll() (*[]SentenceHistoryModel, error) {
	db := databases.GetDB()
	var history []SentenceHistoryModel
	if err := db.Preload("User").Preload("Sentence").Find(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
func UpdateSentenceHistory(id uint, userID uint, sentenceID uint, gameID string, correctness bool) error {
	db := databases.GetDB()
	sentenceHistory, err := GetSentenceHistoryByID(id)
	if err != nil {
		return err
	}
	if sentenceHistory != nil {
		history := &SentenceHistoryModel{
			ID:          id,
			UserID:      userID,
			SentenceID:  sentenceID,
			GameID:      gameID,
			Correctness: correctness,
		}
		return db.Save(history).Error
	}
	return nil
}

func DeleteSentenceHistory(history *SentenceHistoryModel) error {
	db := databases.GetDB()
	return db.Delete(history).Error
}
