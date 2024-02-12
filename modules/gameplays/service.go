package gameplays

import (
	"cp23kk1/common/databases"
	passageRepo "cp23kk1/modules/repository/passage"
	sentenceRepo "cp23kk1/modules/repository/sentence"
	vocabularyRepo "cp23kk1/modules/repository/vocabulary"
)

type VocabularyService struct {
}

func getVocabularies() ([]vocabularyRepo.VocabularyModel, error) {
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())
	return vocabularyRepository.FindManyVocabulary()
}
func randomFromGamePlay() ([]vocabularyRepo.VocabularyModel, error) {
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())

	return vocabularyRepository.RandomVacabulary(50)
}

func randomSentenceForGamePlay() ([]sentenceRepo.SentenceModel, error) {
	vocabularyRepository := sentenceRepo.NewSentenceRepository(databases.GetDB())

	return vocabularyRepository.RandomSentence(50)
}

func randomPassageForGamePlay() ([]passageRepo.PassageModel, error) {
	vocabularyRepository := passageRepo.NewPassageRepository(databases.GetDB())

	return vocabularyRepository.RandomPassage(50)
}
