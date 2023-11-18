package passage

import (
	"cp23kk1/common/databases"
)

func CreatePassage(title string) error {
	db := databases.GetDB()
	passage := &PassageModel{
		Title: title}
	err := db.Create(passage).Error
	return err
}

func FindOnePassage(id int) (*PassageModel, error) {
	db := databases.GetDB()
	var passage PassageModel
	err := db.First(&passage, id).Error
	return &passage, err
}

func FindAllPassages() ([]PassageModel, error) {
	db := databases.GetDB()
	var passages []PassageModel
	err := db.Find(&passages).Error
	return passages, err
}

func UpdatePassage(id int, title string) error {
	db := databases.GetDB()
	passage, err := FindOnePassage(id)
	if err != nil {
		return err
	}
	passage.Title = title
	err = db.Save(passage).Error
	return err
}

func DeletePassage(id int) error {
	db := databases.GetDB()
	passage, err := FindOnePassage(id)
	if err != nil {
		return err
	}

	return db.Delete(passage).Error

}
