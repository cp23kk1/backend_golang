package score

import (
	"net/http"

	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

func AddScoreRoutes(rg *gin.RouterGroup) {
	score := rg.Group("/score")

	// score.Use(auth.AuthMiddleware(true, "access_token"))
	score.GET("/scoreboard", RandomVocabularyForGamePlay)
}

func RandomVocabularyForGamePlay(c *gin.Context) {
	scoreboards, err := getHighScoreBoard()
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "WeeklyScore NotFound", Status: "error"}, map[string]interface{}{}))
		return
	}
	serializer := ScoresSerealizer{c, scoreboards}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Get WeeklyScore successfully", Status: "success"}, map[string]interface{}{"weeklyScore": serializer.Response()}))
}
