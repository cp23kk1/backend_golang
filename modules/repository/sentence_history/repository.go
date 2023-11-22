package sentence_history

import (
	"gorm.io/gorm"
)

type SentenceHistoryRepository struct {
	db *gorm.DB
}

func NewSentenceHistoryRepository(db *gorm.DB) SentenceHistoryRepository {
	return SentenceHistoryRepository{db: db}
}

func (sh SentenceHistoryRepository) CreateSentenceHistory(userID uint, sentenceID uint, gameID string, correctness bool) error {

	history := &SentenceHistoryModel{
		UserID:      userID,
		SentenceID:  sentenceID,
		GameID:      gameID,
		Correctness: correctness,
	}

	return sh.db.Create(history).Error
}
func (sh SentenceHistoryRepository) CreateSentenceHistoryWithArray(userID uint, sentences []SentenceFromGameResultModel, gameID string) error {
	if len(sentences) == 0 {
		return nil
	}
	history := []*SentenceHistoryModel{}

	for _, s := range sentences {
		history = append(history, &SentenceHistoryModel{UserID: userID, SentenceID: uint(s.SentenceID), Correctness: s.Correctness, GameID: gameID})
	}
	return sh.db.Create(history).Error
}
func (sh SentenceHistoryRepository) FindSentenceHistoryByID(id uint) (*SentenceHistoryModel, error) {
	var history SentenceHistoryModel
	if err := sh.db.Where("id = ?", id).Preload("User").Preload("Sentence").First(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func (sh SentenceHistoryRepository) FindSentenceHistoryAll() (*[]SentenceHistoryModel, error) {
	var history []SentenceHistoryModel
	if err := sh.db.Preload("User").Preload("Sentence").Find(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
func (sh SentenceHistoryRepository) UpdateSentenceHistory(id, userID uint, sentenceID uint, gameID string, correctness bool) error {
	sentenceHistory, err := sh.FindSentenceHistoryByID(id)
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
		return sh.db.Save(history).Error
	}
	return nil
}

func (sh SentenceHistoryRepository) DeleteSentenceHistory(history *SentenceHistoryModel) error {
	return sh.db.Delete(history).Error
}
