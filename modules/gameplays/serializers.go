package gameplays

import (
	"cp23kk1/common/databases"
	"cp23kk1/common/enum"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type VocabSerealizer struct {
	C *gin.Context
	databases.VocabularyModel
}
type VocabsSerealizer struct {
	C      *gin.Context
	vocabs []databases.VocabularyModel
}

type VocabResponse struct {
	ID             string `json:"id"`
	Vocabulary     string `json:"vocabulary"`
	Meaning        string `json:"meaning"`
	Pos            string `json:"pos"`
	DifficultyCefr string `json:"difficulty"`
}

func (self *VocabSerealizer) Response() VocabResponse {
	vocabularyModel := self.VocabularyModel
	var vocabs VocabResponse
	err := mapstructure.Decode(vocabularyModel, &vocabs)
	if err != nil {
		fmt.Println("Error mapping data:", err)
		return vocabs
	}
	return vocabs
}

func (self *VocabsSerealizer) Response() []VocabResponse {
	vocabularyModel := self.vocabs
	var vocabs []VocabResponse

	for _, vocab := range vocabularyModel {
		serializer := VocabSerealizer{self.C, vocab}
		vocabs = append(vocabs, serializer.Response())

	}
	return vocabs
}

type SentenceSerealizer struct {
	C *gin.Context
	databases.SentenceModel
}
type SentencesSerealizer struct {
	C         *gin.Context
	sentences []databases.SentenceModel
}

type SentenceResponse struct {
	ID      string `json:"id"`
	Text    string `json:"text"`
	Meaning string `json:"meaning"`
}

func (self *SentenceSerealizer) Response() SentenceResponse {
	sentenceModel := self.SentenceModel
	var sentence SentenceResponse
	err := mapstructure.Decode(sentenceModel, &sentence)
	if err != nil {
		fmt.Println("Error mapping data:", err)
		return sentence
	}
	return sentence
}

func (self *SentencesSerealizer) Response() []SentenceResponse {
	sentenceModel := self.sentences
	var sentences []SentenceResponse

	for _, sentence := range sentenceModel {
		serializer := SentenceSerealizer{self.C, sentence}
		sentences = append(sentences, serializer.Response())

	}
	return sentences
}

type PassageSerealizer struct {
	C *gin.Context
	databases.PassageModel
}
type PassagesSerealizer struct {
	C        *gin.Context
	passages []databases.PassageModel
}

type PassageResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (self *PassageSerealizer) Response() PassageResponse {
	passageModel := self.PassageModel
	var passage PassageResponse
	err := mapstructure.Decode(passageModel, &passage)
	if err != nil {
		fmt.Println("Error mapping data:", err)
		return passage
	}
	return passage
}

func (self *PassagesSerealizer) Response() []PassageResponse {
	passageModel := self.passages
	var passages []PassageResponse

	for _, passage := range passageModel {
		serializer := PassageSerealizer{self.C, passage}
		passages = append(passages, serializer.Response())

	}
	return passages
}

type QuestionSerealizer struct {
	C *gin.Context
	QuestionModel
}
type QuestionsSerealizer struct {
	C         *gin.Context
	questions []QuestionModel
}

type QuestionModel struct {
	DataID       string            `json:"dataId"`
	Question     string            `json:"question"`
	Answers      []AnswerModel     `json:"answers"`
	Pos          *string           `json:"pos"`
	QuestionType enum.QuestionType `json:"questionsType"`
}
type QuestionPassageModel struct {
	DataID       string            `json:"dataId"`
	Questions    []QuestionModel   `json:"questions"`
	Title        string            `json:"title"`
	QuestionType enum.QuestionType `json:"questionType"`
}

type AnswerModel struct {
	Answer      string `json:"answer"`
	Correctness bool   `json:"correctness"`
}
type QuestionResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func (self *QuestionSerealizer) Response() QuestionResponse {
	questionModel := self.QuestionModel
	var question QuestionResponse
	err := mapstructure.Decode(questionModel, &question)
	if err != nil {
		fmt.Println("Error mapping data:", err)
		return question
	}
	return question
}

func (self *QuestionsSerealizer) Response() []QuestionResponse {
	questionModel := self.questions
	var questions []QuestionResponse

	for _, question := range questionModel {
		serializer := QuestionSerealizer{self.C, question}
		questions = append(questions, serializer.Response())
	}
	return questions
}
