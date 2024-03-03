package gameplays

import (
	"cp23kk1/common/databases"
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
	ID             int    `json:"id"`
	Word           string `json:"word"`
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
	ID      uint   `json:"id"`
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
	ID    uint   `json:"id"`
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
