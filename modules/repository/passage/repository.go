package passage

import (
	"cp23kk1/common/databases"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&PassageModel{})
}

func (PassageModel) TableName() string {
	return "passage"
}

func CreatePassage(title string) {
	db := databases.GetDB()
	passage := &PassageModel{
		Title: title}
	db.Create(passage)
}

func FindOnePassage(id int) *PassageModel {
	db := databases.GetDB()
	var passage PassageModel
	db.First(&passage, id)
	if passage.ID == 0 {
		return nil // Record not found
	}
	return &passage
}

func FindAllPassages() []PassageModel {
	db := databases.GetDB()
	var passages []PassageModel
	db.Find(&passages)
	return passages
}

func UpdatePassage(id int, title string) {
	db := databases.GetDB()
	passage := FindOnePassage(id)
	if passage != nil {
		passage.Title = title
		db.Save(passage)
	}
}

func DeletePassage(id int) {
	db := databases.GetDB()
	passage := FindOnePassage(id)
	if passage != nil {
		db.Delete(passage)
	}
}
