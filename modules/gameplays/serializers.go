package gameplays

import (
	sentenceRepo "cp23kk1/modules/repository/sentence"
	vocabularyRepo "cp23kk1/modules/repository/vocabulary"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type VocabSerealizer struct {
	C *gin.Context
	vocabularyRepo.VocabularyModel
}
type VocabsSerealizer struct {
	C      *gin.Context
	vocabs []vocabularyRepo.VocabularyModel
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
	sentenceRepo.SentenceModel
}
type SentencesSerealizer struct {
	C         *gin.Context
	sentences []sentenceRepo.SentenceModel
}

type SentenceResponse struct {
	ID      uint   `json:"id"`
	Text    string `json:"text"`
	Meaning string `json:"meaning"`
}

func (self *SentenceSerealizer) Response() VocabResponse {
	sentenceModel := self.SentenceModel
	var vocabs VocabResponse
	err := mapstructure.Decode(sentenceModel, &vocabs)
	if err != nil {
		fmt.Println("Error mapping data:", err)
		return vocabs
	}
	return vocabs
}

func (self *SentencesSerealizer) Response() []VocabResponse {
	sentenceModel := self.sentences
	var sentences []VocabResponse

	for _, sentence := range sentenceModel {
		serializer := SentenceSerealizer{self.C, sentence}
		sentences = append(sentences, serializer.Response())

	}
	return sentences
}
