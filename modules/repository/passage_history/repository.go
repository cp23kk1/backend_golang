package passage_history

import (
	"cp23kk1/common/databases"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {

	db.AutoMigrate(&PassageHistoryModel{})
}

func (PassageHistoryModel) TableName() string {
	return "passage_history"
}

func CreatePassageHistory(userID, passageID int, gameID string, correctness bool) error {
	db := databases.GetDB()

	passageHistory := PassageHistoryModel{
		UserID:      userID,
		PassageID:   passageID,
		GameID:      gameID,
		Correctness: correctness,
	}
	err := db.Create(&passageHistory).Error
	return err
}

func FindAllPassages() []PassageHistoryModel {
	db := databases.GetDB()
	var passages []PassageHistoryModel
	db.Find(&passages)
	return passages
}
func FindPassageHistoryByID(id int) (*PassageHistoryModel, error) {
	db := databases.GetDB()

	var passageHistory PassageHistoryModel
	if result := db.Preload("User").Preload("Passage").First(&passageHistory, id); result.Error != nil {
		return nil, result.Error
	}
	return &passageHistory, nil
}

func FindPassageHistoriesByUserID(userID int) ([]PassageHistoryModel, error) {
	db := databases.GetDB()

	var passageHistories []PassageHistoryModel
	if result := db.Where("user_id = ?", userID).Preload("User").Preload("Passage").Find(&passageHistories); result.Error != nil {
		return nil, result.Error
	}
	return passageHistories, nil
}

func DeletePassageHistory(id int) {
	db := databases.GetDB()

	passageHistory, _ := FindPassageHistoryByID(id)
	if passageHistory != nil {
		db.Delete(passageHistory)
	}
}
