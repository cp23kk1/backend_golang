package vocabulary_history

import (
	"gorm.io/gorm"
)

type VocabularyHistoryRepository struct {
	db *gorm.DB
}

func NewVocabularyHistoryRepository(db *gorm.DB) VocabularyHistoryRepository {
	return VocabularyHistoryRepository{db: db}
}

func (vh VocabularyHistoryRepository) CreateVocabularyHistory(userID uint, vocabularyID uint, gameID string, correctness bool) error {

	history := &VocabularyHistoryModel{
		UserID:       userID,
		VocabularyID: vocabularyID,
		GameID:       gameID,
		Correctness:  correctness,
	}

	return vh.db.Create(history).Error
}
func (vh VocabularyHistoryRepository) CreateVocabularyHistoryWithArray(userID uint, vocabularies []VocabularyFromGameResultModel, gameID string) error {
	if len(vocabularies) == 0 {
		return nil
	}
	history := []*VocabularyHistoryModel{}

	for _, v := range vocabularies {
		history = append(history, &VocabularyHistoryModel{UserID: userID, VocabularyID: uint(v.VocabularyID), Correctness: v.Correctness, GameID: gameID})
	}
	return vh.db.Create(history).Error
}
func (vh VocabularyHistoryRepository) FindVocabularyHistoryByID(id uint) (*VocabularyHistoryModel, error) {

	var history VocabularyHistoryModel
	if err := vh.db.Where("id = ?", id).Preload("User").Preload("Vocabulary").First(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func (vh VocabularyHistoryRepository) FindVocabularyHistoryAll() (*[]VocabularyHistoryModel, error) {

	var history []VocabularyHistoryModel
	if err := vh.db.Preload("User").Preload("Vocabulary").Find(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
func (vh VocabularyHistoryRepository) UpdateVocabularyHistory(id, userID, vocabularyID uint, gameID string, correctness bool) error {
	vocabularyHistory, err := vh.FindVocabularyHistoryByID(id)
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
		return vh.db.Save(history).Error
	}
	return nil
}

func (vh VocabularyHistoryRepository) DeleteVocabularyHistory(history *VocabularyHistoryModel) error {

	return vh.db.Delete(history).Error
}
