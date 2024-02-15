package history

import (
	"cp23kk1/common"
	"cp23kk1/modules/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHistoryRoutes(rg *gin.RouterGroup) {
	history := rg.Group("/history")

	// history.GET("/vocabulary", VocabulariesHistoryRetrieve)
	// history.GET("/sentence", SentencesHistoryRetrieve)
	// history.GET("/passage", PassagesHistoryRetrieve)
	// history.POST("/game-result", GameResult)
	history.POST("/passage-history", CreatePassageHistory)
	history.Use(auth.AuthMiddleware(true, "access_token"))
	history.POST("/game-result", GameResult)

}

func CreatePassageHistory(c *gin.Context) {

}

func GameResult(c *gin.Context) {
	gameResultValidator := NewGameResultModelValidator()
	if err := gameResultValidator.Bind(c); err != nil {

		c.JSON(http.StatusBadRequest, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "error"}, map[string]interface{}{"errorMessage": err.Error()}))
		return
	}
	userId := int(c.MustGet("userId").(float64))
	fmt.Println("userId", userId)

	if err := gameResult(gameResultValidator, userId); err != nil {
		c.JSON(http.StatusBadRequest, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "error"}, map[string]interface{}{"errorMessage": err.Error()}))
		return
	}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{}))
}
