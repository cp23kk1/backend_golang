package passage_history

import (
	"cp23kk1/common/databases"

	"gorm.io/gorm"
)

type PassageHistoryRepository struct {
	db *gorm.DB
}

func NewPassageHistoryRepository(db *gorm.DB) PassageHistoryRepository {
	return PassageHistoryRepository{db: db}
}

func (ph PassageHistoryRepository) CreatePassageHistory(userID, passageID uint, gameID string, correctness bool) error {

	passageHistory := databases.PassageHistoryModel{
		UserID:      userID,
		PassageID:   passageID,
		GameID:      gameID,
		Correctness: correctness,
	}
	err := ph.db.Create(&passageHistory).Error
	return err
}
func (ph PassageHistoryRepository) CreatePassageHistoryWithArray(userID uint, passages []PassageFromGameResultModel, gameID string) error {
	if len(passages) == 0 {
		return nil
	}
	history := []*databases.PassageHistoryModel{}

	for _, p := range passages {
		history = append(history, &databases.PassageHistoryModel{UserID: userID, PassageID: uint(p.PassageID), Correctness: p.Correctness, GameID: gameID})
	}
	return ph.db.Create(history).Error
}

func (ph PassageHistoryRepository) FindAllPassagesHistory() ([]databases.PassageHistoryModel, error) {

	var passages []databases.PassageHistoryModel
	err := ph.db.Preload("User").Preload("Passage").Find(&passages).Error
	return passages, err
}
func (ph PassageHistoryRepository) FindPassageHistoryByID(id int) (*databases.PassageHistoryModel, error) {

	var passageHistory databases.PassageHistoryModel
	if result := ph.db.Preload("User").Preload("Passage").First(&passageHistory, id); result.Error != nil {
		return nil, result.Error
	}
	return &passageHistory, nil
}

func (ph PassageHistoryRepository) FindPassageHistoriesByUserID(userID int) ([]databases.PassageHistoryModel, error) {

	var passageHistories []databases.PassageHistoryModel
	if result := ph.db.Where("user_id = ?", userID).Preload("User").Preload("Passage").Find(&passageHistories); result.Error != nil {
		return nil, result.Error
	}
	return passageHistories, nil
}

func (ph PassageHistoryRepository) DeletePassageHistory(history *databases.PassageHistoryModel) error {

	return ph.db.Delete(history).Error
}
