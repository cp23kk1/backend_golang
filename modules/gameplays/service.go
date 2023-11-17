package gameplays

import (
	vocabularyRepo "cp23kk1/modules/repository/vocabulary"
)

type VocabularyService struct{}

func getVocabularies() ([]vocabularyRepo.VocabularyModel, error) {
	return vocabularyRepo.FindManyVocabulary()
}
func randomFromGamePlay() ([]vocabularyRepo.VocabularyModel, error) {
	return vocabularyRepo.RandomVacabulary(50)
}
