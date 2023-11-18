// vocabulary_history/routes.go

package vocabulary_history

import (
	// Import common utilities or middleware if needed
	vocabularyHistoryRepo "cp23kk1/modules/repository/vocabulary_history"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupVocabularyHistoryRoutes(router *gin.RouterGroup) {

	vocabularyHistoryGroup := router.Group("/vocabulary-history")

	vocabularyHistoryGroup.POST("/", CreateVocabularyHistoryHandler)
	vocabularyHistoryGroup.GET("/:id", GetVocabularyHistoryHandler)
	vocabularyHistoryGroup.GET("/", GetAllVocabularyHistoryHandler)
	vocabularyHistoryGroup.PUT("/:id", UpdateVocabularyHistoryHandler)
	vocabularyHistoryGroup.DELETE("/:id", DeleteVocabularyHistoryHandler)
}
func CreateVocabularyHistoryHandler(c *gin.Context) {
	vocabularyHistoryModelValidator := NewVocabularyHistoryModelValidator()
	if err := vocabularyHistoryModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := vocabularyHistoryRepo.CreateVocabularyHistory(vocabularyHistoryModelValidator.UserID, vocabularyHistoryModelValidator.VocabularyID, vocabularyHistoryModelValidator.GameID, vocabularyHistoryModelValidator.Correctness); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create VocabularyHistory"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "VocabularyHistory created"})

}

// GetVocabularyHistoryHandler retrieves a single VocabularyHistory record by its ID.
func GetVocabularyHistoryHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	history, err := vocabularyHistoryRepo.FindVocabularyHistoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "VocabularyHistory not found"})
		return
	}

	c.JSON(http.StatusOK, history)
}

// GetAllVocabularyHistoryHandler retrieves all VocabularyHistory records.
func GetAllVocabularyHistoryHandler(c *gin.Context) {
	histories, err := vocabularyHistoryRepo.FindVocabularyHistoryAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vocabulary History Find All error"})
		return
	}

	c.JSON(http.StatusOK, histories)
}

// UpdateVocabularyHistoryHandler updates an existing VocabularyHistory record.
func UpdateVocabularyHistoryHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	vocabularyHistoryModelValidator := NewVocabularyHistoryModelValidator()
	if err := vocabularyHistoryModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := vocabularyHistoryRepo.UpdateVocabularyHistory(id, vocabularyHistoryModelValidator.UserID, vocabularyHistoryModelValidator.VocabularyID, vocabularyHistoryModelValidator.GameID, vocabularyHistoryModelValidator.Correctness); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update VocabularyHistory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "VocabularyHistory updated"})
}

// DeleteVocabularyHistoryHandler deletes a VocabularyHistory record by its ID.
func DeleteVocabularyHistoryHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	history, err := vocabularyHistoryRepo.FindVocabularyHistoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "VocabularyHistory not found"})
		return
	}

	if err := vocabularyHistoryRepo.DeleteVocabularyHistory(history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete VocabularyHistory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "VocabularyHistory deleted"})
}
