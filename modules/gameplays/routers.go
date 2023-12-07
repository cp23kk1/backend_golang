package gameplays

import (
	"net/http"

	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

func AddGameplayRoutes(rg *gin.RouterGroup) {
	gameplay := rg.Group("/gameplays")

	gameplay.GET("/vocabulary", RandomVocabularyForGamePlay)
	gameplay.POST("/game-result", GameResult)
}

func VocabulariesRetrieve(c *gin.Context) {
	vocabs, err := getVocabularies()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("eiei", err)) // need to change later
		return
	}
	c.JSON(http.StatusOK, gin.H{"vocabs": vocabs})
}

func RandomVocabularyForGamePlay(c *gin.Context) {
	vocabs, err := randomFromGamePlay()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("eiei", err)) // need to change later
		return
	}
	serializer := VocabsSerealizer{c, vocabs}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"vocabs": serializer.Response()}))
}

func GameResult(c *gin.Context) {
	gameResultValidator := NewGameResultModelValidator()
	if err := gameResultValidator.Bind(c); err != nil {

		c.JSON(http.StatusBadRequest, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "error"}, map[string]interface{}{"errorMessage": err.Error()}))
		return
	}
	if err := gameResult(gameResultValidator); err != nil {

		c.JSON(http.StatusBadRequest, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "error"}, map[string]interface{}{"errorMessage": err.Error()}))
		return
	}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{}))
}
