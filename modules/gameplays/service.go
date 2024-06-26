package gameplays

import (
	"cp23kk1/common/databases"
	"cp23kk1/common/enum"
	passageRepo "cp23kk1/modules/repository/passage"
	sentenceRepo "cp23kk1/modules/repository/sentence"
	vocabularyRepo "cp23kk1/modules/repository/vocabulary"
	"math/rand"
	"time"
)

func getVocabularies() ([]databases.VocabularyModel, error) {
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())
	return vocabularyRepository.FindManyVocabulary()
}
func randomFromGamePlay() ([]databases.VocabularyModel, error) {
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())

	return vocabularyRepository.RandomVacabulary(50)
}

func randomSentenceForGamePlay() ([]databases.SentenceModel, error) {
	sentenceRepository := sentenceRepo.NewSentenceRepository(databases.GetDB())

	return sentenceRepository.RandomSentence(50)
}

func randomPassageForGamePlay() ([]databases.PassageModel, error) {
	passageRepository := passageRepo.NewPassageRepository(databases.GetDB())

	return passageRepository.RandomPassage(50)
}

func randomQuestionForMultiPlayerGameplay(mode string, numberOfQuestion int) ([]QuestionModel, error) {
	result := []QuestionModel{}

	db := databases.GetDB()
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(db)
	// sentenceRepository := sentenceRepo.NewSentenceRepository(db)
	// passageRepository := passageRepo.NewPassageRepository(db)

	vocabs, err := vocabularyRepository.RandomVacabulary(numberOfQuestion)
	if err != nil {
		return []QuestionModel{}, err
	}
	// sentences, err := sentenceRepository.RandomSentence(5)
	// if err != nil {
	// 	return []QuestionModel{}, QuestionPassageModel{}, err
	// }
	// passage, err := passageRepository.RandomPassage(1)
	// if err != nil {
	// 	return []QuestionModel{}, QuestionPassageModel{}, err
	// }
	for vocabIndex := range vocabs {
		vocabsForAnswer, _ := vocabularyRepository.FindManyVocabularyNotSameVocabByPosAndLimit(vocabs[vocabIndex].ID, vocabs[vocabIndex].POS, 2)
		answerVocabs := mapVocabToAnswer(vocabsForAnswer)
		answerVocabs = append(answerVocabs, AnswerModel{AnswerID: vocabs[vocabIndex].ID, Answer: vocabs[vocabIndex].Meaning, Correctness: true})
		result = append(result, QuestionModel{
			Question:        vocabs[vocabIndex].Vocabulary,
			Pos:             &vocabs[vocabIndex].POS,
			Answers:         answerVocabs,
			QuestionType:    enum.VOCABULARY,
			CorrectAnswerID: vocabs[vocabIndex].ID,
			DataID:          vocabs[vocabIndex].ID})
	}
	// result = generatedSentenceQuestion(sentences, result, vocabularyRepository, enum.SENTENCE)

	// passageQuestion := QuestionPassageModel{DataID: passage[0].ID,
	// 	Questions: generatedSentenceQuestion(passage[0].Sentences, []QuestionModel{}, vocabularyRepository, enum.SENTENCE),
	// 	Title:     passage[0].Title, QuestionType: enum.PASSAGE}

	return result, nil
}
func randomQuestionForGameplay() ([]QuestionModel, QuestionPassageModel, error) {
	result := []QuestionModel{}

	db := databases.GetDB()
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(db)
	sentenceRepository := sentenceRepo.NewSentenceRepository(db)
	passageRepository := passageRepo.NewPassageRepository(db)

	vocabs, err := vocabularyRepository.RandomVacabulary(5)
	if err != nil {
		return []QuestionModel{}, QuestionPassageModel{}, err
	}
	sentences, err := sentenceRepository.RandomSentence(5)
	if err != nil {
		return []QuestionModel{}, QuestionPassageModel{}, err
	}
	passage, err := passageRepository.RandomPassage(1)
	if err != nil {
		return []QuestionModel{}, QuestionPassageModel{}, err
	}
	for vocabIndex := range vocabs {
		vocabsForAnswer, _ := vocabularyRepository.FindManyVocabularyNotSameVocabByPosAndLimit(vocabs[vocabIndex].ID, vocabs[vocabIndex].POS, 2)
		answerVocabs := mapVocabToAnswer(vocabsForAnswer)
		answerVocabs = append(answerVocabs, AnswerModel{AnswerID: vocabs[vocabIndex].ID, Answer: vocabs[vocabIndex].Meaning, Correctness: true})
		result = append(result, QuestionModel{
			Question:        vocabs[vocabIndex].Vocabulary,
			Pos:             &vocabs[vocabIndex].POS,
			Answers:         answerVocabs,
			QuestionType:    enum.VOCABULARY,
			CorrectAnswerID: vocabs[vocabIndex].ID,
			DataID:          vocabs[vocabIndex].ID})
	}
	result = generatedSentenceQuestion(sentences, result, vocabularyRepository, enum.SENTENCE, 1)

	passageQuestion := QuestionPassageModel{DataID: passage[0].ID,
		Questions: generatedSentenceQuestion(passage[0].Sentences, []QuestionModel{}, vocabularyRepository, enum.SENTENCE, 0),
		Title:     passage[0].Title, QuestionType: enum.PASSAGE}

	return result, passageQuestion, nil
}

func mapVocabToAnswer(vocabs []databases.VocabularyModel) []AnswerModel {
	answerVocabs := []AnswerModel{}
	for vocabIndex := range vocabs {
		answerVocabs = append(answerVocabs, AnswerModel{AnswerID: vocabs[vocabIndex].ID, Answer: vocabs[vocabIndex].Meaning, Correctness: false})
	}
	return answerVocabs
}

func mapVocabToSentenceAnswer(vocabs []databases.VocabularyModel) []AnswerModel {
	answerVocabs := []AnswerModel{}
	for vocabIndex := range vocabs {
		answerVocabs = append(answerVocabs, AnswerModel{AnswerID: vocabs[vocabIndex].ID, Answer: vocabs[vocabIndex].Vocabulary, Correctness: false})
	}
	return answerVocabs
}

func generatedSentenceQuestion(sentences []databases.SentenceModel, result []QuestionModel, vocabularyRepository vocabularyRepo.VocabularyRepository, questionType enum.QuestionType, numAnswer int) []QuestionModel {
	rand.Seed(time.Now().UnixNano())

	temp := result
	for sentenceIndex := range sentences {
		lengthSentenceAnswers := len(sentences[sentenceIndex].Vocabularies)
		if lengthSentenceAnswers == 0 {
			continue
		}
		randIndex := rand.Intn(lengthSentenceAnswers)
		sentenceCorrectAnswer := sentences[sentenceIndex].Vocabularies[randIndex]
		sentenceOtherAnswer, _ := vocabularyRepository.FindManyVocabularyNotSameVocabByPosAndLimit(sentenceCorrectAnswer.ID, sentenceCorrectAnswer.POS, numAnswer)
		answerSentence := mapVocabToSentenceAnswer(sentenceOtherAnswer)
		answerSentence = append(answerSentence, AnswerModel{AnswerID: sentenceCorrectAnswer.ID, Answer: sentenceCorrectAnswer.Vocabulary,
			Correctness: true})

		temp = append(temp, QuestionModel{
			Question:        sentences[sentenceIndex].Sentence,
			Answers:         answerSentence,
			Pos:             &sentences[sentenceIndex].Tense,
			QuestionType:    questionType,
			CorrectAnswerID: sentences[sentenceIndex].Vocabularies[randIndex].ID,
			DataID:          sentences[sentenceIndex].ID})
	}
	return temp
}
