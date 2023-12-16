// vocabulary_history/routes.go

package vocabulary_history

import (
	// Import common utilities or middleware if needed
	"cp23kk1/common/databases"
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
	vhRepository := vocabularyHistoryRepo.NewVocabularyHistoryRepository(databases.GetDB())
	if err := vhRepository.CreateVocabularyHistory(uint(vocabularyHistoryModelValidator.UserID), uint(vocabularyHistoryModelValidator.VocabularyID), vocabularyHistoryModelValidator.GameID, vocabularyHistoryModelValidator.Correctness); err != nil {
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
	vhRepository := vocabularyHistoryRepo.NewVocabularyHistoryRepository(databases.GetDB())

	history, err := vhRepository.FindVocabularyHistoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "VocabularyHistory not found"})
		return
	}

	c.JSON(http.StatusOK, history)
}

// GetAllVocabularyHistoryHandler retrieves all VocabularyHistory records.
func GetAllVocabularyHistoryHandler(c *gin.Context) {
	vhRepository := vocabularyHistoryRepo.NewVocabularyHistoryRepository(databases.GetDB())

	histories, err := vhRepository.FindVocabularyHistoryAll()
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
	vhRepository := vocabularyHistoryRepo.NewVocabularyHistoryRepository(databases.GetDB())

	if err := vhRepository.UpdateVocabularyHistory(uint(id), uint(vocabularyHistoryModelValidator.UserID), uint(vocabularyHistoryModelValidator.VocabularyID), vocabularyHistoryModelValidator.GameID, vocabularyHistoryModelValidator.Correctness); err != nil {
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
	vhRepository := vocabularyHistoryRepo.NewVocabularyHistoryRepository(databases.GetDB())

	history, err := vhRepository.FindVocabularyHistoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "VocabularyHistory not found"})
		return
	}

	if err := vhRepository.DeleteVocabularyHistory(history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete VocabularyHistory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "VocabularyHistory deleted"})
}
