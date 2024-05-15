package gameplays

import (
	"net/http"

	"cp23kk1/common"
	"cp23kk1/modules/auth"

	"github.com/gin-gonic/gin"
)

func AddGameplayRoutes(rg *gin.RouterGroup) {
	gameplay := rg.Group("/gameplays")

	gameplay.Use(auth.AuthMiddleware(true, "access_token"))
	gameplay.GET("/vocabulary", RandomVocabularyForGamePlay)
	gameplay.GET("/sentence", RandomSentenceForGamePlay)
	gameplay.GET("/passage", RandomPassageForGamePlay)
	gameplay.GET("/single-player", RandomForSinglePlayer)
	gameplay.POST("/multi-player", RandomForMultiPlayer)
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
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Vocabulary NotFound", Status: "error"}, map[string]interface{}{}))
		return
	}
	serializer := VocabsSerealizer{c, vocabs}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Get Vocabulary successfully", Status: "success"}, map[string]interface{}{"vocabs": serializer.Response()}))
}
func RandomSentenceForGamePlay(c *gin.Context) {
	sentences, err := randomSentenceForGamePlay()
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Sentence NotFound", Status: "error"}, map[string]interface{}{}))
		return
	}
	serializer := SentencesSerealizer{c, sentences}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Get Sentence successfully", Status: "success"}, map[string]interface{}{"sentences": serializer.Response()}))
}
func RandomPassageForGamePlay(c *gin.Context) {
	passages, err := randomPassageForGamePlay()
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Passage NotFound", Status: "error"}, map[string]interface{}{}))
		return
	}
	serializer := PassagesSerealizer{c, passages}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Get Passage successfully", Status: "success"}, map[string]interface{}{"passages": serializer.Response()}))
}
func RandomForSinglePlayer(c *gin.Context) {
	questions, passageQuestion, err := randomQuestionForGameplay()
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Questions NotFound", Status: "error"}, map[string]interface{}{}))
		return
	}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Get Questions successfully", Status: "success"}, map[string]interface{}{"questions": questions, "passageQuestion": passageQuestion}))
}
func RandomForMultiPlayer(c *gin.Context) {
	multiPlayerValidator := NewMultiPlayerValidator()
	if err := multiPlayerValidator.Bind(c); err != nil {

		c.JSON(http.StatusBadRequest, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: err.Error(), Status: "error"}, map[string]interface{}{"errorMessage": err.Error()}))
		return
	}
	questions, err := randomQuestionForMultiPlayerGameplay(multiPlayerValidator.Mode, multiPlayerValidator.NumberOfQuestion)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Questions NotFound", Status: "error"}, map[string]interface{}{}))
		return
	}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Get Questions successfully", Status: "success"}, map[string]interface{}{"questions": questions}))
}
