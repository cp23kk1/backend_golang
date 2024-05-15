package sentence_history

import (
	// Import common utilities or middleware if needed
	"cp23kk1/common/databases"
	sentenceHistoryRepo "cp23kk1/modules/repository/sentence_history"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupSentenceHistoryRoutes(router *gin.RouterGroup) {

	sentenceHistoryGroup := router.Group("/sentence-history")

	sentenceHistoryGroup.POST("/", CreateSentenceHistoryHandler)
	sentenceHistoryGroup.GET("/:id", GetSentenceHistoryHandler)
	sentenceHistoryGroup.GET("/", GetAllSentenceHistoryHandler)
	sentenceHistoryGroup.PUT("/:id", UpdateSentenceHistoryHandler)
	sentenceHistoryGroup.DELETE("/:id", DeleteSentenceHistoryHandler)
}
func CreateSentenceHistoryHandler(c *gin.Context) {
	sentenceHistoryModelValidator := NewSentenceHistoryModelValidator()
	if err := sentenceHistoryModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shRepository := sentenceHistoryRepo.NewSentenceHistoryRepository(databases.GetDB())
	if err := shRepository.CreateSentenceHistory(uint(sentenceHistoryModelValidator.UserID), sentenceHistoryModelValidator.SentenceID, sentenceHistoryModelValidator.GameID, sentenceHistoryModelValidator.Correctness); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create SentenceHistory"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "SentenceHistory created"})

}

// GetSentenceHistoryHandler retrieves a single SentenceHistory record by its ID.
func GetSentenceHistoryHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	shRepository := sentenceHistoryRepo.NewSentenceHistoryRepository(databases.GetDB())
	history, err := shRepository.FindSentenceHistoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SentenceHistory not found"})
		return
	}

	c.JSON(http.StatusOK, history)
}

// GetAllSentenceHistoryHandler retrieves all SentenceHistory records.
func GetAllSentenceHistoryHandler(c *gin.Context) {
	shRepository := sentenceHistoryRepo.NewSentenceHistoryRepository(databases.GetDB())

	histories, _ := shRepository.FindSentenceHistoryAll()

	c.JSON(http.StatusOK, histories)
}

// UpdateSentenceHistoryHandler updates an existing SentenceHistory record.
func UpdateSentenceHistoryHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sentenceHistoryModelValidator := NewSentenceHistoryModelValidator()
	if err := sentenceHistoryModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shRepository := sentenceHistoryRepo.NewSentenceHistoryRepository(databases.GetDB())

	if err := shRepository.UpdateSentenceHistory(uint(id), uint(sentenceHistoryModelValidator.UserID), sentenceHistoryModelValidator.SentenceID, sentenceHistoryModelValidator.GameID, sentenceHistoryModelValidator.Correctness); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update SentenceHistory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SentenceHistory updated"})
}

// DeleteSentenceHistoryHandler deletes a SentenceHistory record by its ID.
func DeleteSentenceHistoryHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	shRepository := sentenceHistoryRepo.NewSentenceHistoryRepository(databases.GetDB())

	history, err := shRepository.FindSentenceHistoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SentenceHistory not found"})
		return
	}

	if err := shRepository.DeleteSentenceHistory(history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete SentenceHistory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SentenceHistory deleted"})
}
