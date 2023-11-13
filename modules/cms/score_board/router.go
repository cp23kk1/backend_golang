package score_board

import (
	ScoreBoardRepo "cp23kk1/modules/repository/score_board"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupScoreBoardRoutes(router *gin.RouterGroup) {
	scoreBoardGroup := router.Group("/score_board")
	{
		scoreBoardGroup.POST("/", CreateScoreBoardHandler)
		scoreBoardGroup.GET("/:id", GetScoreBoardByIDHandler)
		scoreBoardGroup.GET("/user/:userID", GetScoreBoardsByUserIDHandler)
		scoreBoardGroup.GET("/", GetAllScoreBoardsHandler)
		scoreBoardGroup.DELETE("/:id", DeleteScoreBoardHandler)
	}
}
func CreateScoreBoardHandler(c *gin.Context) {
	// Parse request data
	scoreBoardModelValidator := NewScoreBoardModelValidator()
	if err := scoreBoardModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	startDate, err := time.Parse("2006-01-02", scoreBoardModelValidator.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	endDate, err := time.Parse("2006-01-02", scoreBoardModelValidator.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create the score board record
	ScoreBoardRepo.CreateScoreBoard(scoreBoardModelValidator.UserID,
		scoreBoardModelValidator.Score,
		scoreBoardModelValidator.Week,
		startDate,
		endDate)
	c.JSON(http.StatusCreated, gin.H{"message": "Score board created successfully"})
}

func GetScoreBoardByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	scoreBoard, err := ScoreBoardRepo.FindScoreBoardByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Score board not found"})
		return
	}

	c.JSON(http.StatusOK, scoreBoard)
}

func GetScoreBoardsByUserIDHandler(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))

	scoreBoards, err := ScoreBoardRepo.FindScoreBoardsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, scoreBoards)
}

func GetAllScoreBoardsHandler(c *gin.Context) {
	scoreBoards, err := ScoreBoardRepo.FindAllScoreBoards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, scoreBoards)
}

func DeleteScoreBoardHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Check if the score board exists
	_, err := ScoreBoardRepo.FindScoreBoardByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Score board not found"})
		return
	}

	// Delete the score board
	ScoreBoardRepo.DeleteScoreBoard(id)
	c.JSON(http.StatusOK, gin.H{"message": "Score board deleted successfully"})
}
