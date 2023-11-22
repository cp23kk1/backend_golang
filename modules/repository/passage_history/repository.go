package passage_history

import (
	"gorm.io/gorm"
)

type PassageHistoryRepository struct {
	db *gorm.DB
}

func NewPassageHistoryRepository(db *gorm.DB) PassageHistoryRepository {
	return PassageHistoryRepository{db: db}
}

func (ph PassageHistoryRepository) CreatePassageHistory(userID, passageID uint, gameID string, correctness bool) error {

	passageHistory := PassageHistoryModel{
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
	history := []*PassageHistoryModel{}

	for _, p := range passages {
		history = append(history, &PassageHistoryModel{UserID: userID, PassageID: uint(p.PassageID), Correctness: p.Correctness, GameID: gameID})
	}
	return ph.db.Create(history).Error
}

func (ph PassageHistoryRepository) FindAllPassagesHistory() ([]PassageHistoryModel, error) {

	var passages []PassageHistoryModel
	err := ph.db.Preload("User").Preload("Passage").Find(&passages).Error
	return passages, err
}
func (ph PassageHistoryRepository) FindPassageHistoryByID(id int) (*PassageHistoryModel, error) {

	var passageHistory PassageHistoryModel
	if result := ph.db.Preload("User").Preload("Passage").First(&passageHistory, id); result.Error != nil {
		return nil, result.Error
	}
	return &passageHistory, nil
}

func (ph PassageHistoryRepository) FindPassageHistoriesByUserID(userID int) ([]PassageHistoryModel, error) {

	var passageHistories []PassageHistoryModel
	if result := ph.db.Where("user_id = ?", userID).Preload("User").Preload("Passage").Find(&passageHistories); result.Error != nil {
		return nil, result.Error
	}
	return passageHistories, nil
}

func (ph PassageHistoryRepository) DeletePassageHistory(history *PassageHistoryModel) error {

	return ph.db.Delete(history).Error
}
