package vocabulary_history

import (
	"gorm.io/gorm"
)

type VocabularyHistoryRepository struct {
	db *gorm.DB
}

func NewVocabularyHistoryRepository(db *gorm.DB) *VocabularyHistoryRepository {
	return &VocabularyHistoryRepository{db}
}

func (r *VocabularyHistoryRepository) CreateVocabularyHistory(history *VocabularyHistory) error {
	return r.db.Create(history).Error
}

func (r *VocabularyHistoryRepository) GetVocabularyHistoryByID(id uint) (*VocabularyHistory, error) {
	var history VocabularyHistory
	if err := r.db.Where("id = ?", id).Preload("User").Preload("Vocabulary").First(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func (r *VocabularyHistoryRepository) UpdateVocabularyHistory(history *VocabularyHistory) error {
	return r.db.Save(history).Error
}

func (r *VocabularyHistoryRepository) DeleteVocabularyHistory(history *VocabularyHistory) error {
	return r.db.Delete(history).Error
}
