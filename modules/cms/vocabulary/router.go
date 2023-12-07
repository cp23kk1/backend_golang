// vocabulary/routes.go

package vocabulary

import (
	"cp23kk1/common/databases"
	vocabularyRepo "cp23kk1/modules/repository/vocabulary"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupVocabularyRoutes(router *gin.RouterGroup) {
	vocabularyGroup := router.Group("/vocabularies")

	vocabularyGroup.POST("/", CreateVocabularyHandler)
	vocabularyGroup.GET("/:id", GetVocabularyHandler)
	vocabularyGroup.GET("/", GetAllVocabulariesHandler)
	vocabularyGroup.PUT("/:id", UpdateVocabularyHandler)
	vocabularyGroup.DELETE("/:id", DeleteVocabularyHandler)
}

// vocabulary/handlers.go
func CreateVocabularyHandler(c *gin.Context) {
	vocabularyModelValidator := NewVocabularyModelValidator()
	if err := vocabularyModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())
	vocabularyRepository.CreateVocabulary(
		vocabularyModelValidator.Word,
		vocabularyModelValidator.Meaning,
		vocabularyModelValidator.Pos,
		vocabularyModelValidator.DifficultyCefr,
	)

	c.JSON(http.StatusCreated, gin.H{"message": "Vocabulary created"})
}

func GetVocabularyHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())

	vocabulary := vocabularyRepository.FindOneVocabulary(id)
	if vocabulary == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vocabulary not found"})
		return
	}

	c.JSON(http.StatusOK, vocabulary)
}

func GetAllVocabulariesHandler(c *gin.Context) {
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())

	vocabularies, err := vocabularyRepository.FindManyVocabulary()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vocabulary Get All Error"})
		return
	}
	c.JSON(http.StatusOK, vocabularies)
}

func UpdateVocabularyHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	vocabularyModelValidator := NewVocabularyModelValidator()
	if err := vocabularyModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())

	vocabularyRepository.UpdateVocabulary(
		id,
		vocabularyModelValidator.Word,
		vocabularyModelValidator.Meaning,
		vocabularyModelValidator.Pos,
		vocabularyModelValidator.DifficultyCefr,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Vocabulary updated"})
}

func DeleteVocabularyHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	vocabularyRepository := vocabularyRepo.NewVocabularyRepository(databases.GetDB())

	vocabularyRepository.DeleteVocabulary(id)
	c.JSON(http.StatusOK, gin.H{"message": "Vocabulary deleted"})
}
