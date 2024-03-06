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

func (v VocabularyRepository) CreateVocabulary(word, pos, tag, lemma, definition, dep string, difficultyCefr uint, meaning string) {
	vocabulary := &databases.VocabularyModel{
		Vocabulary:   word,
		Meaning:      meaning,
		POS:          pos,
		DifficultyID: difficultyCefr,
		Definition:   definition,
		Tag:          tag,
		Lemma:        lemma,
		Dep:          dep,
	}
	v.db.Create(vocabulary)
}

func (v VocabularyRepository) FindOneVocabulary(id string) *databases.VocabularyModel {
	var vocabulary databases.VocabularyModel
	err := v.db.First(&vocabulary, id).Error
	if err != nil {
		return nil // Record not found
	}
	return &vocabulary
}

func (v VocabularyRepository) FindManyVocabulary() ([]databases.VocabularyModel, error) {
	var vocabularies []databases.VocabularyModel
	err := v.db.Find(&vocabularies).Error
	return vocabularies, err
}

func (v VocabularyRepository) FindManyVocabularyNotSameVocabByPosAndLimit(vocabularyId string, pos string, limit int) ([]databases.VocabularyModel, error) {
	var vocabularies []databases.VocabularyModel
	err := v.db.Where("id <> ?", vocabularyId).Where("pos = ?", pos).Limit(limit).Find(&vocabularies).Error
	return vocabularies, err
}

func (v VocabularyRepository) UpdateVocabulary(id string, word, pos, tag, lemma, definition, dep string, difficultyCefr uint, meaning string) {
	vocabulary := v.FindOneVocabulary(id)
	if vocabulary != nil {
		vocabulary.Vocabulary = word
		vocabulary.Meaning = meaning
		vocabulary.POS = pos
		vocabulary.DifficultyID = difficultyCefr
		vocabulary.Meaning = meaning
		vocabulary.Definition = definition
		vocabulary.Tag = tag
		vocabulary.Lemma = lemma
		vocabulary.Dep = dep
		v.db.Save(vocabulary)
	}
}

func (v VocabularyRepository) DeleteVocabulary(id string) {
	vocabulary := v.FindOneVocabulary(id)
	if vocabulary != nil {
		v.db.Delete(vocabulary)
	}
}

func (v VocabularyRepository) RandomVacabulary(limit int) ([]databases.VocabularyModel, error) {

	var vocabularies []databases.VocabularyModel
	// v.db.Model(&ScoreBoardModel{}).Preload("User").Find(&scoreBoards).Error
	err := v.db.Preload("Sentences").Order("RAND()").Limit(limit).Find(&vocabularies).Error
	return vocabularies, err
}
