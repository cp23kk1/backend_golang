package score

import (
	"net/http"

	"cp23kk1/common"
	"cp23kk1/modules/auth"

	"github.com/gin-gonic/gin"
)

func AddScoreRoutes(rg *gin.RouterGroup) {
	score := rg.Group("/score")

	score.Use(auth.AuthMiddleware(true, "access_token"))
	score.GET("/scoreboard", getScoreBoard)
	score.GET("/bestscore", getBestScore)
}

func getScoreBoard(c *gin.Context) {
	scoreBoards, err := getHighScoreBoard()
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "WeeklyScore NotFound", Status: "error"}, map[string]interface{}{}))
		return
	}
	userScoreBoard, _ := getUserScoreBoard(c.MustGet("userId").(uint))
	serializer := ScoresSerealizer{c, scoreBoards}
	serializer1 := ScoreSerealizer{c, userScoreBoard}

	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Get WeeklyScore successfully", Status: "success"}, map[string]interface{}{"weeklyScore": serializer.Response(), "userScore": serializer1.Response()}))
}

func getBestScore(c *gin.Context) {
	userId := c.MustGet("userId").(uint)
	bestScore, err := getBestScoreUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "BestScore NotFound", Status: "error"}, map[string]interface{}{}))
		return
	}
	serializer := ScoresSerealizer{c, bestScore}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Get BestScore successfully", Status: "success"}, map[string]interface{}{"bestScore": serializer.BestScoreResponse()}))
}
