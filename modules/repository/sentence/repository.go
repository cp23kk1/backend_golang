package sentence

import (
	"cp23kk1/common/databases"

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
	sentence := &databases.SentenceModel{
		PassageID: passageId,
		Sequence:  sequence,
		Text:      text,
		Meaning:   meaning,
	}
	return s.db.Create(sentence).Error
}

// GetSentenceByID retrieves a sentence record from the database by its ID.
func (s SentenceRepository) FindSentenceByID(id int) (*databases.SentenceModel, error) {
	var sentence databases.SentenceModel
	err := s.db.Preload("Passage").First(&sentence, id).Error
	return &sentence, err
}
func (s SentenceRepository) FindAllSentence() (*[]databases.SentenceModel, error) {
	var sentence []databases.SentenceModel
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
	return s.db.Delete(&databases.SentenceModel{}, id).Error
}

func (v SentenceRepository) RandomSentence(limit int) ([]databases.SentenceModel, error) {

	var sentences []databases.SentenceModel
	// v.db.Model(&ScoreBoardModel{}).Preload("User").Find(&scoreBoards).Error
	err := v.db.Model(&databases.SentenceModel{}).Order("RAND()").Limit(limit).Scan(&sentences).Error
	return sentences, err
}
