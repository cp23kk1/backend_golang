package score_board

import (
	"cp23kk1/common/databases"
	"time"
)

func CreateScoreBoard(userID, score, week int, startDate, endDate time.Time) error {
	db := databases.GetDB()

	scoreBoard := ScoreBoardModel{
		UserID:    userID,
		Score:     score,
		Week:      week,
		StartDate: startDate,
		EndDate:   endDate,
	}
	return db.Create(&scoreBoard).Error
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
func DeleteScoreBoard(id int) error {
	db := databases.GetDB()

	scoreBoard, err := FindScoreBoardByID(id)
	if err != nil {
		return err
	}

	return db.Delete(scoreBoard).Error

}
