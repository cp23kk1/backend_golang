package vocabulary_history

import "gorm.io/gorm"

func CreateVocabularyHistory(db *gorm.DB, vh *VocabularyHistory) error {
	result := db.Create(vh)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetVocabularyHistoryByID(db *gorm.DB, id uint) (*VocabularyHistory, error) {
	var vh VocabularyHistory
	result := db.First(&vh, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &vh, nil
}

func UpdateVocabularyHistory(db *gorm.DB, vh *VocabularyHistory) error {
	result := db.Save(vh)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteVocabularyHistory(db *gorm.DB, vh *VocabularyHistory) error {
	result := db.Delete(vh)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
