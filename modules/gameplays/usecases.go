package gameplays

type VocabularyUseCase struct{}

func getVocabularies() ([]VocabularyModel, error) {
	return selectAll()
}
