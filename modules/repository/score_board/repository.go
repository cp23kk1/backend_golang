package score_board

import (
	"cp23kk1/common/databases"
	"time"

	"gorm.io/gorm"
)

type ScoreBoardRepository struct {
	db *gorm.DB
}

func NewScoreBoardRepository(db *gorm.DB) ScoreBoardRepository {
	return ScoreBoardRepository{db: db}
}

func (s *ScoreBoardRepository) CreateScoreBoard(userID uint, score, week int, startDate, endDate time.Time) error {

	scoreBoard := databases.ScoreBoardModel{
		UserID:    userID,
		Score:     score,
		Week:      week,
		StartDate: startDate,
		EndDate:   endDate,
	}
	return s.db.Create(&scoreBoard).Error
}
func (s *ScoreBoardRepository) FindScoreBoardByID(id int) (*databases.ScoreBoardModel, error) {

	var scoreBoard databases.ScoreBoardModel
	if result := s.db.Preload("User").First(&scoreBoard, id); result.Error != nil {
		return nil, result.Error
	}
	return &scoreBoard, nil
}
func (s *ScoreBoardRepository) FindScoreBoardsByUserID(userID int) ([]databases.ScoreBoardModel, error) {

	var scoreBoards []databases.ScoreBoardModel
	if result := s.db.Where("user_id = ?", userID).Preload("User").Find(&scoreBoards); result.Error != nil {
		return nil, result.Error
	}
	return scoreBoards, nil
}
func (s *ScoreBoardRepository) FindAllScoreBoards() ([]databases.ScoreBoardModel, error) {

	var scoreBoards []databases.ScoreBoardModel
	err := s.db.Model(&databases.ScoreBoardModel{}).Preload("User").Find(&scoreBoards).Error

	return scoreBoards, err
}

func (s *ScoreBoardRepository) FindAllHighScoreBoardsByWeekLimit(limit, week int) ([]databases.ScoreBoardModel, error) {

	var scoreBoards []databases.ScoreBoardModel
	err := s.db.Model(&databases.ScoreBoardModel{}).Preload("User").Where("week = ?", week).Order("score desc").Limit(limit).Find(&scoreBoards).Error

	return scoreBoards, err
}
func (s *ScoreBoardRepository) DeleteScoreBoard(id int) error {

	scoreBoard, err := s.FindScoreBoardByID(id)
	if err != nil {
		return err
	}

	return s.db.Delete(scoreBoard).Error

}
