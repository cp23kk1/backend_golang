package gameplays

import (
	"cp23kk1/common/databases"
	"cp23kk1/modules/repository/passage_history"
	"cp23kk1/modules/repository/score_board"
	"cp23kk1/modules/repository/sentence_history"
	vocabularyRepo "cp23kk1/modules/repository/vocabulary"
	"cp23kk1/modules/repository/vocabulary_history"
	"time"
)

type VocabularyService struct{}

func getVocabularies() ([]vocabularyRepo.VocabularyModel, error) {
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())
	return vocabularyRepository.FindManyVocabulary()
}
func randomFromGamePlay() ([]vocabularyRepo.VocabularyModel, error) {
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())

	return vocabularyRepository.RandomVacabulary(50)
}
func gameResult(gameResultValidator GameResultModelValidator) error {
	tx := databases.GetDB().Begin()
	userId := 2
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
