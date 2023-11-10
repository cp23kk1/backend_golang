package score_board

import (
	"cp23kk1/common/databases"
	"time"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {

	db.AutoMigrate(&ScoreBoardModel{})
}

func (ScoreBoardModel) TableName() string {
	return "score_board"
}

func CreateScoreBoard(userID, score, week int, startDate, endDate time.Time) {
	db := databases.GetDB()

	scoreBoard := ScoreBoardModel{
		UserID:    userID,
		Score:     score,
		Week:      week,
		StartDate: startDate,
		EndDate:   endDate,
	}
	db.Create(&scoreBoard)
}
func FindScoreBoardByID(id int) (*ScoreBoardModel, error) {
	db := databases.GetDB()
	var scoreBoard ScoreBoardModel
	if result := db.Preload("User").First(&scoreBoard, id); result.Error != nil {
		return nil, result.Error
	}
	return &scoreBoard, nil
}
func FindScoreBoardsByUserID(userID int) ([]ScoreBoardModel, error) {
	db := databases.GetDB()

	var scoreBoards []ScoreBoardModel
	if result := db.Where("user_id = ?", userID).Preload("User").Find(&scoreBoards); result.Error != nil {
		return nil, result.Error
	}
	return scoreBoards, nil
}
func FindAllScoreBoards() ([]ScoreBoardModel, error) {
	db := databases.GetDB()

	var scoreBoards []ScoreBoardModel
	err := db.Model(&ScoreBoardModel{}).Preload("User").Find(&scoreBoards).Error

	return scoreBoards, err
}
func DeleteScoreBoard(id int) {
	db := databases.GetDB()

	scoreBoard, _ := FindScoreBoardByID(id)
	if scoreBoard != nil {
		db.Delete(scoreBoard)
	}
}
