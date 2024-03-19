package score_board

import (
	"cp23kk1/common/databases"
	ScoreBoardRepo "cp23kk1/modules/repository/score_board"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupScoreBoardRoutes(router *gin.RouterGroup) {
	scoreBoardGroup := router.Group("/score-board")
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
	scoreBoardRepository := ScoreBoardRepo.NewScoreBoardRepository(databases.GetDB())
	// Create the score board record
	scoreBoardRepository.CreateScoreBoard(uint(scoreBoardModelValidator.UserID),
		scoreBoardModelValidator.Score,
		scoreBoardModelValidator.Week,
		startDate,
		endDate,
		scoreBoardModelValidator.GameID, scoreBoardModelValidator.Mode)
	c.JSON(http.StatusCreated, gin.H{"message": "Score board created successfully"})
}

func GetScoreBoardByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	scoreBoardRepository := ScoreBoardRepo.NewScoreBoardRepository(databases.GetDB())

	scoreBoard, err := scoreBoardRepository.FindScoreBoardByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Score board not found"})
		return
	}

	c.JSON(http.StatusOK, scoreBoard)
}

func GetScoreBoardsByUserIDHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	scoreBoardRepository := ScoreBoardRepo.NewScoreBoardRepository(databases.GetDB())

	scoreBoards, err := scoreBoardRepository.FindScoreBoardsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, scoreBoards)
}

func GetAllScoreBoardsHandler(c *gin.Context) {
	scoreBoardRepository := ScoreBoardRepo.NewScoreBoardRepository(databases.GetDB())

	scoreBoards, err := scoreBoardRepository.FindAllScoreBoards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, scoreBoards)
}

func DeleteScoreBoardHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	scoreBoardRepository := ScoreBoardRepo.NewScoreBoardRepository(databases.GetDB())

	// Check if the score board exists
	_, err := scoreBoardRepository.FindScoreBoardByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Score board not found"})
		return
	}

	// Delete the score board
	scoreBoardRepository.DeleteScoreBoard(id)
	c.JSON(http.StatusOK, gin.H{"message": "Score board deleted successfully"})
}
