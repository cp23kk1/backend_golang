package passage

import (
	"cp23kk1/common/databases"

	"gorm.io/gorm"
)

type PassageRepository struct {
	db *gorm.DB
}

func NewPassageRepository(db *gorm.DB) PassageRepository {

	return PassageRepository{db: db}
}
func (p PassageRepository) CreatePassage(title string) error {
	passage := &databases.PassageModel{
		Title: title}
	err := p.db.Create(passage).Error
	return err
}

func (p PassageRepository) FindOnePassage(id int) (*databases.PassageModel, error) {
	var passage databases.PassageModel
	err := p.db.First(&passage, id).Error
	return &passage, err
}

func (p PassageRepository) FindAllPassages() ([]databases.PassageModel, error) {
	var passages []databases.PassageModel
	err := p.db.Find(&passages).Error
	return passages, err
}

func (p PassageRepository) UpdatePassage(id int, title string) error {
	passage, err := p.FindOnePassage(id)
	if err != nil {
		return err
	}
	passage.Title = title
	err = p.db.Save(passage).Error
	return err
}

func (p PassageRepository) DeletePassage(id int) error {
	passage, err := p.FindOnePassage(id)
	if err != nil {
		return err
	}

	return p.db.Delete(passage).Error

}
func (v PassageRepository) RandomPassage(limit int) ([]databases.PassageModel, error) {

	var passages []databases.PassageModel
	err := v.db.Preload("Sentences").Preload("Sentences.Vocabularies").Order("RAND()").Limit(limit).Find(&passages).Error
	return passages, err
}
