package history

import (
	"cp23kk1/common/databases"
	"cp23kk1/modules/repository/passage_history"
	passageHistoryRepo "cp23kk1/modules/repository/passage_history"
	"cp23kk1/modules/repository/score_board"
	scoreBoardRepo "cp23kk1/modules/repository/score_board"
	"cp23kk1/modules/repository/sentence_history"
	"cp23kk1/modules/repository/vocabulary_history"
	"time"
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

func gameResult(gameResultValidator GameResultModelValidator) error {
	tx := databases.GetDB().Begin()

	userId := 2
	if gameResultValidator.UserID != nil {
		userId = *gameResultValidator.UserID
	}
	scoreBoardRepository := score_board.NewScoreBoardRepository(tx)
	vhRepository := vocabulary_history.NewVocabularyHistoryRepository(tx)
	shRepository := sentence_history.NewSentenceHistoryRepository(tx)
	phRepository := passage_history.NewPassageHistoryRepository(tx)
	if err := scoreBoardRepository.CreateScoreBoard(uint(userId), gameResultValidator.CurrentSocore, 1, time.Now(), time.Now()); err != nil {
		tx.Rollback()
		return err
	}
	if err := vhRepository.CreateVocabularyHistoryWithArray(uint(userId), gameResultValidator.Vocabs, gameResultValidator.GameID); err != nil {
		tx.Rollback()
		return err
	}
	if err := shRepository.CreateSentenceHistoryWithArray(uint(userId), gameResultValidator.Sentences, gameResultValidator.GameID); err != nil {
		tx.Rollback()
		return err
	}
	if err := phRepository.CreatePassageHistoryWithArray(uint(userId), gameResultValidator.Passages, gameResultValidator.GameID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
