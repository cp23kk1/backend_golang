package history

import (
	"cp23kk1/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHistoryRoutes(rg *gin.RouterGroup) {
	history := rg.Group("/history")

	// history.GET("/vocabulary", VocabulariesHistoryRetrieve)
	// history.GET("/sentence", SentencesHistoryRetrieve)
	// history.GET("/passage", PassagesHistoryRetrieve)
	// history.POST("/game-result", GameResult)
	history.GET("/game-result", GameResult)

	history.POST("/passage-history", CreatePassageHistory)

}

func GameResult(c *gin.Context) {
	result, _ := getScoreBoard()
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"vocabs": result}))
}

func CreatePassageHistory(c *gin.Context) {

}
