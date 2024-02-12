package passage

import (
	"gorm.io/gorm"
)

type PassageRepository struct {
	db *gorm.DB
}

func NewPassageRepository(db *gorm.DB) PassageRepository {

	return PassageRepository{db: db}
}
func (p PassageRepository) CreatePassage(title string) error {
	passage := &PassageModel{
		Title: title}
	err := p.db.Create(passage).Error
	return err
}

func (p PassageRepository) FindOnePassage(id int) (*PassageModel, error) {
	var passage PassageModel
	err := p.db.First(&passage, id).Error
	return &passage, err
}

func (p PassageRepository) FindAllPassages() ([]PassageModel, error) {
	var passages []PassageModel
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
func (v PassageRepository) RandomPassage(limit int) ([]PassageModel, error) {

	var passages []PassageModel
	err := v.db.Model(&PassageModel{}).Order("RAND()").Limit(limit).Scan(&passages).Error
	return passages, err
}
