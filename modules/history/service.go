package history

import (
	"cp23kk1/common/databases"
	passageHistoryRepo "cp23kk1/modules/repository/passage_history"
	scoreBoardRepo "cp23kk1/modules/repository/score_board"
)

func createPassageHistory(userID, passageID uint, gameID string, correctness bool) error {
	phRepository := passageHistoryRepo.NewPassageHistoryRepository(databases.GetDB())
	err := phRepository.CreatePassageHistory(userID, passageID, gameID, correctness)
	return err
}
func getPassageHistory() ([]passageHistoryRepo.PassageHistoryModel, error) {
	phRepository := passageHistoryRepo.NewPassageHistoryRepository(databases.GetDB())

	return phRepository.FindAllPassagesHistory()
}

func getScoreBoard() ([]scoreBoardRepo.ScoreBoardModel, error) {
	scoreBoardRepository := scoreBoardRepo.NewScoreBoardRepository(databases.GetDB())

	return scoreBoardRepository.FindAllScoreBoards()
}
