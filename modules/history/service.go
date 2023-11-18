package history

import (
	passageHistoryRepo "cp23kk1/modules/repository/passage_history"
	scoreBoardRepo "cp23kk1/modules/repository/score_board"
)

func createPassageHistory(userID, passageID int, gameID string, correctness bool) error {
	err := passageHistoryRepo.CreatePassageHistory(userID, passageID, gameID, correctness)
	return err
}
func getPassageHistory() ([]passageHistoryRepo.PassageHistoryModel, error) {
	return passageHistoryRepo.FindAllPassagesHistory()
}

func getScoreBoard() ([]scoreBoardRepo.ScoreBoardModel, error) {
	return scoreBoardRepo.FindAllScoreBoards()
}
