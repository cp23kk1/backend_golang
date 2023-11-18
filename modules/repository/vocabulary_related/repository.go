package vocabulary_related

import (
	"cp23kk1/common/databases"
)

// CreateVocabularyRelated creates a new vocabulary related record in the database
func CreateVocabularyRelated(vocabularyID, sentenceID int) error {
	db := databases.GetDB()
	vocabularyRelated := &VocabularyRelatedModel{
		VocabularyID: vocabularyID,
		SentenceID:   sentenceID,
	}
	return db.Create(vocabularyRelated).Error
}

// GetVocabularyRelatedByID retrieves a vocabulary related record by ID from the database
func FindVocabularyRelatedByID(vocabularyID, sentenceID int) (*VocabularyRelatedModel, error) {
	db := databases.GetDB()
	var vocabularyRelated VocabularyRelatedModel
	err := db.Where("vocabulary_id = ? AND sentence_id = ?", vocabularyID, sentenceID).Preload("Vocabulary").Preload("Sentence").First(&vocabularyRelated).Error
	if err != nil {
		return nil, err
	}
	return &vocabularyRelated, nil
}

// UpdateVocabularyRelated updates a vocabulary related record in the database
func UpdateVocabularyRelated(vocabularyRelated *VocabularyRelatedModel) error {
	db := databases.GetDB()
	return db.Save(vocabularyRelated).Error
}

// DeleteVocabularyRelated deletes a vocabulary related record from the database
func DeleteVocabularyRelated(vocabularyID, sentenceID int) error {
	db := databases.GetDB()
	return db.Where("vocabulary_id = ? AND sentence_id = ?", vocabularyID, sentenceID).Delete(&VocabularyRelatedModel{}).Error
}

// GetAllVocabularyRelated retrieves all VocabularyRelated records from the database
func FindAllVocabularyRelated() ([]VocabularyRelatedModel, error) {
	db := databases.GetDB()

	var vocabularyRelatedList []VocabularyRelatedModel
	err := db.Preload("Vocabulary").Preload("Sentence").Find(&vocabularyRelatedList).Error
	if err != nil {
		return nil, err
	}
	return vocabularyRelatedList, nil
}
