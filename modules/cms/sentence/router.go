package sentence

import (
	SentenceRepo "cp23kk1/modules/repository/sentence"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupSentenceRoutes(router *gin.RouterGroup) {
	sentenceGroup := router.Group("/sentence")
	{
		sentenceGroup.POST("/", CreateSentenceHandler)
		sentenceGroup.GET("/:id", GetSentenceByIDHandler)
		sentenceGroup.PUT("/:id", UpdateSentenceByIDHandler)
		// sentenceGroup.GET("/user/:userID", GetSentencesByUserIDHandler)
		sentenceGroup.GET("/", GetAllSentencesHandler)
		sentenceGroup.DELETE("/:id", DeleteSentenceHandler)
	}
}

func CreateSentenceHandler(c *gin.Context) {
	// Parse request data
	sentenceModelValidator := NewSentenceModelValidator()
	if err := sentenceModelValidator.Bind(c); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	// Create the sentence record
	err := SentenceRepo.CreateSentence(
		sentenceModelValidator.PassageID,
		sentenceModelValidator.Sequence,
		sentenceModelValidator.Text,
		sentenceModelValidator.Meaning,
	)

	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to create sentence", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Sentence created successfully"})
}

func UpdateSentenceByIDHandler(c *gin.Context) {
	// Parse request data
	id, _ := strconv.Atoi(c.Param("id"))
	sentenceModelValidator := NewSentenceModelValidator()
	if err := sentenceModelValidator.Bind(c); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	// Create the sentence record
	err := SentenceRepo.UpdateSentence(
		id,
		sentenceModelValidator.PassageID,
		sentenceModelValidator.Sequence,
		sentenceModelValidator.Text,
		sentenceModelValidator.Meaning,
	)

	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to update sentence", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Sentence update successfully"})
}

func GetSentenceByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	sentence, err := SentenceRepo.FindSentenceByID(id)
	if err != nil {
		handleError(c, http.StatusNotFound, "Sentence not found", err)
		return
	}

	c.JSON(http.StatusOK, sentence)
}

// func GetSentencesByUserIDHandler(c *gin.Context) {
// 	userID, _ := strconv.Atoi(c.Param("userID"))

// 	sentences, err := SentenceRepo.FindSentencesByUserID(userID)
// 	if err != nil {
// 		handleError(c, http.StatusInternalServerError, "Internal server error", err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, sentences)
// }

func GetAllSentencesHandler(c *gin.Context) {
	sentences, err := SentenceRepo.FindAllSentence()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, sentences)
}

func DeleteSentenceHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Check if the sentence exists
	_, err := SentenceRepo.FindSentenceByID(id)
	if err != nil {
		handleError(c, http.StatusNotFound, "Sentence not found", err)
		return
	}

	// Delete the sentence
	err = SentenceRepo.DeleteSentence(id)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to delete sentence", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sentence deleted successfully"})
}

func handleError(c *gin.Context, statusCode int, message string, err error) {
	c.JSON(statusCode, gin.H{"error": message, "details": err.Error()})
}
