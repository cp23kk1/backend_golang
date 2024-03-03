package vocabulary

import (
	"cp23kk1/common/databases"

	"gorm.io/gorm"
)

type VocabularyRepository struct {
	db *gorm.DB
}

func NewVocabularyRepository(db *gorm.DB) VocabularyRepository {
	return VocabularyRepository{db: db}
}

func (v VocabularyRepository) CreateVocabulary(word, pos, difficultyCefr string, meaning string) {
	vocabulary := &databases.VocabularyModel{
		Word:           word,
		Meaning:        meaning,
		POS:            pos,
		DifficultyCEFR: difficultyCefr,
	}
	v.db.Create(vocabulary)
}

func (v VocabularyRepository) FindOneVocabulary(id int) *databases.VocabularyModel {
	var vocabulary databases.VocabularyModel
	v.db.First(&vocabulary, id)
	if vocabulary.ID == 0 {
		return nil // Record not found
	}
	return &vocabulary
}

func (v VocabularyRepository) FindManyVocabulary() ([]databases.VocabularyModel, error) {
	var vocabularies []databases.VocabularyModel
	err := v.db.Find(&vocabularies).Error
	return vocabularies, err
}

func (v VocabularyRepository) UpdateVocabulary(id int, word, pos, difficultyCefr string, meaning string) {
	vocabulary := v.FindOneVocabulary(id)
	if vocabulary != nil {
		vocabulary.Word = word
		vocabulary.Meaning = meaning
		vocabulary.POS = pos
		vocabulary.DifficultyCEFR = difficultyCefr
		v.db.Save(vocabulary)
	}
}

func (v VocabularyRepository) DeleteVocabulary(id int) {
	vocabulary := v.FindOneVocabulary(id)
	if vocabulary != nil {
		v.db.Delete(vocabulary)
	}
}

func (v VocabularyRepository) RandomVacabulary(limit int) ([]databases.VocabularyModel, error) {

	var vocabularies []databases.VocabularyModel
	// v.db.Model(&ScoreBoardModel{}).Preload("User").Find(&scoreBoards).Error
	err := v.db.Model(&databases.VocabularyModel{}).Order("RAND()").Limit(limit).Scan(&vocabularies).Error
	return vocabularies, err
}
