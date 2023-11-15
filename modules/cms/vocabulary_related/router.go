package vocabulary_related

import (
	vocabularyRelatedRepo "cp23kk1/modules/repository/vocabulary_related"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupVocabularyRelatedRoutes(router *gin.RouterGroup) {
	vocabularyRelatedGroup := router.Group("/vocabulary-related")

	vocabularyRelatedGroup.POST("/", CreateVocabularyRelatedHandler)
	vocabularyRelatedGroup.GET("/", GetAllVocabularyRelatedHandler)
	vocabularyRelatedGroup.GET("/:vocabularyID/:sentenceID", GetVocabularyRelatedHandler)
	// vocabularyRelatedGroup.PUT("/", UpdateVocabularyRelatedHandler)
	vocabularyRelatedGroup.DELETE("/:vocabularyID/:sentenceID", DeleteVocabularyRelatedHandler)
}

func CreateVocabularyRelatedHandler(c *gin.Context) {
	vocabularyRelatedModelValidator := NewVocabularyRelatedModelValidator()
	if err := vocabularyRelatedModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := vocabularyRelatedRepo.CreateVocabularyRelated(vocabularyRelatedModelValidator.VocabularyID, vocabularyRelatedModelValidator.SentenceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Model"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "VocabularyRelated created"})
}

func GetVocabularyRelatedHandler(c *gin.Context) {
	vocabularyID, err := strconv.Atoi(c.Param("vocabularyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vocabularyID"})
		return
	}

	sentenceID, err := strconv.Atoi(c.Param("sentenceID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sentenceID"})
		return
	}

	vocabularyRelated, err := vocabularyRelatedRepo.GetVocabularyRelatedByID(vocabularyID, sentenceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "VocabularyRelated not found"})
		return
	}

	c.JSON(http.StatusOK, vocabularyRelated)
}
func GetAllVocabularyRelatedHandler(c *gin.Context) {

	vocabularyRelated, err := vocabularyRelatedRepo.GetAllVocabularyRelated()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "VocabularyRelated not found"})
		return
	}

	c.JSON(http.StatusOK, vocabularyRelated)
}
func UpdateVocabularyRelatedHandler(c *gin.Context) {
	var vocabularyRelated vocabularyRelatedRepo.VocabularyRelatedModel
	if err := c.ShouldBindJSON(&vocabularyRelated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := vocabularyRelatedRepo.UpdateVocabularyRelated(&vocabularyRelated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update VocabularyRelated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "VocabularyRelated updated"})
}

func DeleteVocabularyRelatedHandler(c *gin.Context) {
	vocabularyID, err := strconv.Atoi(c.Param("vocabularyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vocabularyID"})
		return
	}

	sentenceID, err := strconv.Atoi(c.Param("sentenceID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sentenceID"})
		return
	}

	err = vocabularyRelatedRepo.DeleteVocabularyRelated(vocabularyID, sentenceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete VocabularyRelated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "VocabularyRelated deleted"})
}
