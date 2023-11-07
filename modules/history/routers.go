package history

import (
	"cp23kk1/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHistoryRoutes(rg *gin.RouterGroup) {
	history := rg.Group("/history")

	history.GET("/vocabulary", VocabulariesHistoryRetrieve)
	history.GET("/sentence", SentencesHistoryRetrieve)
	history.GET("/passage", PassagesHistoryRetrieve)
	history.POST("/game-result", GameResult)
}

func VocabulariesHistoryRetrieve(c *gin.Context) {
	vocabulary_history, err := selectVocabularyHistoryAll()
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "data not found"}, nil))
		return
	}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"vocabulary_history": vocabulary_history}))
}

func SentencesHistoryRetrieve(c *gin.Context) {
	sentence_history, err := selectSentenceHistoryAll()
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "data not found"}, nil))
		return
	}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"sentence_history": sentence_history}))
}
func PassagesHistoryRetrieve(c *gin.Context) {
	passage_history, err := selectPassageHistoryAll()
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "data not found"}, nil))
		return
	}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"passage_history": passage_history}))
}

func GameResult(c *gin.Context) {
	var requestData GameResultRequest

	// Parse the JSON body into a map
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Print(requestData)
	var resultVocabs interface{}
	if len(requestData.Vocabs) > 0 {
		var err error
		resultVocabs, err = insertVocabulary(requestData.Vocabs, requestData.GameID)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Can't insert"}, map[string]interface{}{"vocabs": "can't insert"}))
			return
		}
	}
	fmt.Print(resultVocabs)
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"vocabs": resultVocabs}))

}

// func RandomVocabularyForGamePlay(c *gin.Context) {
// 	vocabs, err := randomFromGamePlay()
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, common.NewError("eiei", err)) // need to change later
// 		return
// 	}
// 	// var vocabularies []VocabularyResponse
// 	// err = mapstructure.Decode(vocabs, &vocabularies)
// 	// if err != nil {
// 	// 	fmt.Println("Error mapping data:", err)
// 	// }
// 	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"vocabs": vocabs}))
// }
