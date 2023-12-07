package score_board

import (
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

	scoreBoard := ScoreBoardModel{
		UserID:    userID,
		Score:     score,
		Week:      week,
		StartDate: startDate,
		EndDate:   endDate,
	}
	return s.db.Create(&scoreBoard).Error
}
func (s *ScoreBoardRepository) FindScoreBoardByID(id int) (*ScoreBoardModel, error) {

	var scoreBoard ScoreBoardModel
	if result := s.db.Preload("User").First(&scoreBoard, id); result.Error != nil {
		return nil, result.Error
	}
	return &scoreBoard, nil
}
func (s *ScoreBoardRepository) FindScoreBoardsByUserID(userID int) ([]ScoreBoardModel, error) {

	var scoreBoards []ScoreBoardModel
	if result := s.db.Where("user_id = ?", userID).Preload("User").Find(&scoreBoards); result.Error != nil {
		return nil, result.Error
	}
	return scoreBoards, nil
}
func (s *ScoreBoardRepository) FindAllScoreBoards() ([]ScoreBoardModel, error) {

	var scoreBoards []ScoreBoardModel
	err := s.db.Model(&ScoreBoardModel{}).Preload("User").Find(&scoreBoards).Error

	return scoreBoards, err
}
func (s *ScoreBoardRepository) DeleteScoreBoard(id int) error {

	scoreBoard, err := s.FindScoreBoardByID(id)
	if err != nil {
		return err
	}

	return s.db.Delete(scoreBoard).Error

}
