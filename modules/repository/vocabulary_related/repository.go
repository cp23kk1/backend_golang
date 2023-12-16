package vocabulary_related

import (
	"gorm.io/gorm"
)

type VocaubularyRelatedRepository struct {
	db *gorm.DB
}

func NewVocaubularyRelated(db *gorm.DB) VocaubularyRelatedRepository {
	return VocaubularyRelatedRepository{db: db}
}

// CreateVocabularyRelated creates a new vocabulary related record in the database
func (vr VocaubularyRelatedRepository) CreateVocabularyRelated(vocabularyID, sentenceID int) error {
	vocabularyRelated := &VocabularyRelatedModel{
		VocabularyID: vocabularyID,
		SentenceID:   sentenceID,
	}
	return vr.db.Create(vocabularyRelated).Error
}

// GetVocabularyRelatedByID retrieves a vocabulary related record by ID from the database
func (vr VocaubularyRelatedRepository) FindVocabularyRelatedByID(vocabularyID, sentenceID int) (*VocabularyRelatedModel, error) {
	var vocabularyRelated VocabularyRelatedModel
	err := vr.db.Where("vocabulary_id = ? AND sentence_id = ?", vocabularyID, sentenceID).Preload("Vocabulary").Preload("Sentence").First(&vocabularyRelated).Error
	if err != nil {
		return nil, err
	}
	return &vocabularyRelated, nil
}

// UpdateVocabularyRelated updates a vocabulary related record in the database
func (vr VocaubularyRelatedRepository) UpdateVocabularyRelated(vocabularyRelated *VocabularyRelatedModel) error {
	return vr.db.Save(vocabularyRelated).Error
}

// DeleteVocabularyRelated deletes a vocabulary related record from the database
func (vr VocaubularyRelatedRepository) DeleteVocabularyRelated(vocabularyID, sentenceID int) error {
	return vr.db.Where("vocabulary_id = ? AND sentence_id = ?", vocabularyID, sentenceID).Delete(&VocabularyRelatedModel{}).Error
}

// GetAllVocabularyRelated retrieves all VocabularyRelated records from the database
func (vr VocaubularyRelatedRepository) FindAllVocabularyRelated() ([]VocabularyRelatedModel, error) {

	var vocabularyRelatedList []VocabularyRelatedModel
	err := vr.db.Preload("Vocabulary").Preload("Sentence").Find(&vocabularyRelatedList).Error
	if err != nil {
		return nil, err
	}
	return vocabularyRelatedList, nil
}
