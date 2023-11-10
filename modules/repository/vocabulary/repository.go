package vocabulary

import (
	"cp23kk1/common/databases"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {

	db.AutoMigrate(&VocabularyModel{})
}

func (VocabularyModel) TableName() string {
	return "vocabulary"
}

func CreateVocabulary(word, pos, difficultyCefr string, meaning string) {
	db := databases.GetDB()
	vocabulary := &VocabularyModel{
		Word:           word,
		Meaning:        meaning,
		Pos:            pos,
		DifficultyCefr: difficultyCefr,
	}
	db.Create(vocabulary)
}

func FindOneVocabulary(id int) *VocabularyModel {
	db := databases.GetDB()
	var vocabulary VocabularyModel
	db.First(&vocabulary, id)
	if vocabulary.ID == 0 {
		return nil // Record not found
	}
	return &vocabulary
}

func FindManyVocabulary() []VocabularyModel {
	db := databases.GetDB()
	var vocabularies []VocabularyModel
	db.Find(&vocabularies)
	return vocabularies
}

func UpdateVocabulary(id int, word, pos, difficultyCefr string, meaning string) {
	db := databases.GetDB()
	vocabulary := FindOneVocabulary(id)
	if vocabulary != nil {
		vocabulary.Word = word
		vocabulary.Meaning = meaning
		vocabulary.Pos = pos
		vocabulary.DifficultyCefr = difficultyCefr
		db.Save(vocabulary)
	}
}

func DeleteVocabulary(id int) {
	db := databases.GetDB()
	vocabulary := FindOneVocabulary(id)
	if vocabulary != nil {
		db.Delete(vocabulary)
	}
}

func RandomVacabulary(limit int) ([]VocabularyModel, error) {

	db := databases.GetDB()
	var vocabularies []VocabularyModel
	// db.Model(&ScoreBoardModel{}).Preload("User").Find(&scoreBoards).Error
	err := db.Model(&VocabularyModel{}).Order("RAND()").Limit(limit).Scan(&vocabularies).Error
	return vocabularies, err
}
