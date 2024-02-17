package score

import (
	"cp23kk1/common/databases"
	scoreBoardRepo "cp23kk1/modules/repository/score_board"
)

func getHighScoreBoard() ([]databases.ScoreBoardModel, error) {
	scoreBoardRepository := scoreBoardRepo.NewScoreBoardRepository(databases.GetDB())
	return scoreBoardRepository.FindAllHighScoreBoardsByWeekLimit(5, 1)
}
