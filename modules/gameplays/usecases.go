package gameplays

type VocabularyUseCase struct{}

func getVocabularies() ([]VocabularyModel, error) {
	return selectAll()
}
func randomFromGamePlay() ([]VocabularyModel, error) {
	return randomVacabulary()
}
