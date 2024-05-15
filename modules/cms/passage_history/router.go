package passage_history

import (
	// Import common utilities or middleware if needed
	"cp23kk1/common/databases"
	passageHistoryRepo "cp23kk1/modules/repository/passage_history"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupPassageHistoryRoutes(router *gin.RouterGroup) {

	passageHistoryGroup := router.Group("/passage-history")

	passageHistoryGroup.POST("/", CreatePassageHistoryHandler)
	passageHistoryGroup.GET("/:id", GetPassageHistoryByIDHandler)
	passageHistoryGroup.GET("/", GetAllPassagesHistoryHandler)
	passageHistoryGroup.DELETE("/:id", DeletePassageHistoryHandler)
}
func CreatePassageHistoryHandler(c *gin.Context) {
	vocabularyHistoryModelValidator := NewPassageHistoryModelValidator()
	if err := vocabularyHistoryModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	phRepository := passageHistoryRepo.NewPassageHistoryRepository(databases.GetDB())

	err := phRepository.CreatePassageHistory(uint(vocabularyHistoryModelValidator.UserID), vocabularyHistoryModelValidator.PassageID, vocabularyHistoryModelValidator.GameID, vocabularyHistoryModelValidator.Correctness)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Passage History created"})
}

// FindAllPassagesHistoryHandler handles the Find All Passage Histories route
func GetAllPassagesHistoryHandler(c *gin.Context) {
	phRepository := passageHistoryRepo.NewPassageHistoryRepository(databases.GetDB())

	passages, err := phRepository.FindAllPassagesHistory()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passage Get All Error"})
		return
	}
	// P
	c.JSON(200, passages)
}

// FindPassageHistoryByIDHandler handles the Find Passage History by ID route
func GetPassageHistoryByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Parse id as an integer here
	phRepository := passageHistoryRepo.NewPassageHistoryRepository(databases.GetDB())

	passage, err := phRepository.FindPassageHistoryByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Passage History not found"})
		return
	}

	c.JSON(200, passage)
}

// DeletePassageHistoryHandler handles the Delete Passage History by ID route
func DeletePassageHistoryHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	phRepository := passageHistoryRepo.NewPassageHistoryRepository(databases.GetDB())

	history, err := phRepository.FindPassageHistoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PassageHistory not found"})
		return
	}
	// Parse id as an integer here
	if err := phRepository.DeletePassageHistory(history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete PassageHistory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PassageHistory deleted"})
}
