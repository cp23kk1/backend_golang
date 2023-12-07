package sentence

import (
	"gorm.io/gorm"
)

type SentenceRepository struct {
	db *gorm.DB
}

func NewSentenceRepository(db *gorm.DB) SentenceRepository {
	return SentenceRepository{db: db}
}

// CreateSentence creates a new sentence record in the database.
func (s SentenceRepository) CreateSentence(passageId *uint, sequence *int, text, meaning string) error {
	sentence := &SentenceModel{
		PassageID: passageId,
		Sequence:  sequence,
		Text:      text,
		Meaning:   meaning,
	}
	return s.db.Create(sentence).Error
}

// GetSentenceByID retrieves a sentence record from the database by its ID.
func (s SentenceRepository) FindSentenceByID(id int) (*SentenceModel, error) {
	var sentence SentenceModel
	err := s.db.Preload("Passage").First(&sentence, id).Error
	return &sentence, err
}
func (s SentenceRepository) FindAllSentence() (*[]SentenceModel, error) {
	var sentence []SentenceModel
	err := s.db.Preload("Passage").Find(&sentence).Error
	return &sentence, err
}

// UpdateSentence updates an existing sentence record in the database.
func (s SentenceRepository) UpdateSentence(id int, passageId *uint, sequence *int, text, meaning string) error {
	sentence, err := s.FindSentenceByID(id)
	if err != nil {
		return err
	}

	sentence.PassageID = passageId
	sentence.Sequence = sequence
	sentence.Text = text
	sentence.Meaning = meaning

	return s.db.Save(sentence).Error

}

// DeleteSentence deletes a sentence record from the database by its ID.
func (s SentenceRepository) DeleteSentence(id int) error {
	return s.db.Delete(&SentenceModel{}, id).Error
}
