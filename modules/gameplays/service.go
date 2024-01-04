package gameplays

import (
	"cp23kk1/common/databases"
	vocabularyRepo "cp23kk1/modules/repository/vocabulary"
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
