package sentence_history

import (
	"cp23kk1/common/databases"

	"gorm.io/gorm"
)

type SentenceHistoryRepository struct {
	db *gorm.DB
}

func NewSentenceHistoryRepository(db *gorm.DB) SentenceHistoryRepository {
	return SentenceHistoryRepository{db: db}
}

func (sh SentenceHistoryRepository) CreateSentenceHistory(userID uint, sentenceID string, gameID string, correctness bool) error {

	history := &databases.SentenceHistoryModel{
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
	history := []*databases.SentenceHistoryModel{}

	for _, s := range sentences {
		history = append(history, &databases.SentenceHistoryModel{UserID: userID, SentenceID: s.SentenceID, Correctness: s.Correctness, GameID: gameID, VocabularyID: s.AnswerID})
	}
	return sh.db.Create(history).Error
}
func (sh SentenceHistoryRepository) FindSentenceHistoryByID(id uint) (*databases.SentenceHistoryModel, error) {
	var history databases.SentenceHistoryModel
	if err := sh.db.Where("id = ?", id).Preload("User").Preload("Sentence").First(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
func (sh SentenceHistoryRepository) FindSentenceHistoriesByUserID(userID int) ([]databases.SentenceHistoryModel, error) {

	var sentenceHistories []databases.SentenceHistoryModel
	// if result := sh.db.Where("user_id = ?", userID).Preload("User").Preload("Sentence").Find(&sentenceHistories); result.Error != nil {
	if result := sh.db.Where("user_id = ?", userID).Find(&sentenceHistories); result.Error != nil {
		return nil, result.Error
	}
	return sentenceHistories, nil
}
func (sh SentenceHistoryRepository) FindSentenceHistoriesByUserIDAndCorrect(userID int) ([]databases.SentenceHistoryModel, error) {

	var sentenceHistories []databases.SentenceHistoryModel
	// if result := sh.db.Where("user_id = ?", userID).Preload("User").Preload("Sentence").Find(&sentenceHistories); result.Error != nil {
	if result := sh.db.Where("user_id = ?", userID).Where("correctness = true").Find(&sentenceHistories); result.Error != nil {
		return nil, result.Error
	}
	return sentenceHistories, nil
}

func (sh SentenceHistoryRepository) FindSentenceHistoryAll() (*[]databases.SentenceHistoryModel, error) {
	var history []databases.SentenceHistoryModel
	if err := sh.db.Preload("User").Preload("Sentence").Find(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
func (sh SentenceHistoryRepository) UpdateSentenceHistory(id, userID uint, sentenceID string, gameID string, correctness bool) error {
	sentenceHistory, err := sh.FindSentenceHistoryByID(id)
	if err != nil {
		return err
	}
	if sentenceHistory != nil {
		history := &databases.SentenceHistoryModel{
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

func (sh SentenceHistoryRepository) DeleteSentenceHistory(history *databases.SentenceHistoryModel) error {
	return sh.db.Delete(history).Error
}
func (sh SentenceHistoryRepository) FindCountSentenceHistoryGroupByPOS(userId int) []SentenceHistoryCountModel {
	var result []SentenceHistoryCountModel
	sh.db.Raw("select count(s.id) as count ,b.tense as tense from vocaverse.sentence_history s join vocaverse.sentence b on s.sentence_id = b.id where s.user_id = ? group by b.tense order by b.tense;", userId).Find(&result)
	return result
}

func (sh SentenceHistoryRepository) FindCountSentenceHistoryGroupByPOSAndCorrect(userId int) []SentenceHistoryCountModel {
	var result []SentenceHistoryCountModel
	sh.db.Raw("select count(s.id) as count ,b.tense as tense from vocaverse.sentence_history s join vocaverse.sentence b on s.sentence_id = b.id where s.correctness and s.user_id = ? group by b.tense order by b.tense", userId).Find(&result)
	return result
}
