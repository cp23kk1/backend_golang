package sentence

import (
	"cp23kk1/common/databases"
)

// CreateSentence creates a new sentence record in the database.
func CreateSentence(passageId, sequence int, text, meaning string) error {
	db := databases.GetDB()
	sentence := &SentenceModel{
		PassageID: passageId,
		Sequence:  sequence,
		Text:      text,
		Meaning:   meaning,
	}
	return db.Create(sentence).Error
}

// GetSentenceByID retrieves a sentence record from the database by its ID.
func GetSentenceByID(id int) (*SentenceModel, error) {
	db := databases.GetDB()
	var sentence SentenceModel
	err := db.Preload("Passage").First(&sentence, id).Error
	return &sentence, err
}
func GetAllSentence() (*[]SentenceModel, error) {
	db := databases.GetDB()
	var sentence []SentenceModel
	err := db.Preload("Passage").Find(&sentence).Error
	return &sentence, err
}

// UpdateSentence updates an existing sentence record in the database.
func UpdateSentence(id, passageId, sequence int, text, meaning string) error {
	db := databases.GetDB()
	sentence, err := GetSentenceByID(id)
	if err != nil {
		return err
	}

	sentence.PassageID = passageId
	sentence.Sequence = sequence
	sentence.Text = text
	sentence.Meaning = meaning

	return db.Save(sentence).Error

}

// DeleteSentence deletes a sentence record from the database by its ID.
func DeleteSentence(id int) error {
	db := databases.GetDB()
	return db.Delete(&SentenceModel{}, id).Error
}
