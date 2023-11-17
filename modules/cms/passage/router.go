package passage

import (
	passageRepo "cp23kk1/modules/repository/passage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupPassageRoutes(router *gin.RouterGroup) {
	passageGroup := router.Group("/passages")

	passageGroup.POST("/", CreatePassageHandler)
	passageGroup.GET("/:id", GetPassageHandler)
	passageGroup.GET("/", GetAllPassagesHandler)
	passageGroup.PUT("/:id", UpdatePassageHandler)
	passageGroup.DELETE("/:id", DeletePassageHandler)
}

func CreatePassageHandler(c *gin.Context) {
	passageModelValidator := NewPassageModelValidator()
	if err := passageModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passageRepo.CreatePassage(passageModelValidator.Title)
	c.JSON(http.StatusCreated, gin.H{"message": "Passage created"})
}

func GetPassageHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	passage, err := passageRepo.FindOnePassage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Passage not found"})
		return
	}

	c.JSON(http.StatusOK, passage)
}

func GetAllPassagesHandler(c *gin.Context) {
	passages, err := passageRepo.FindAllPassages()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passage Get All Error"})
		return
	}
	c.JSON(http.StatusOK, passages)
}

func UpdatePassageHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	passageModelValidator := NewPassageModelValidator()
	if err := passageModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passageRepo.UpdatePassage(id, passageModelValidator.Title)
	c.JSON(http.StatusOK, gin.H{"message": "Passage updated"})
}

func DeletePassageHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	passageRepo.DeletePassage(id)
	c.JSON(http.StatusOK, gin.H{"message": "Passage deleted"})
}
