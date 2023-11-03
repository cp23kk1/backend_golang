package gameplays

import (
	"errors"
	"fmt"
	"net/http"

	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

func AddGameplayRoutes(rg *gin.RouterGroup) {
	gameplay := rg.Group("/gameplay")

	gameplay.GET("/vocabulary", VocabulariesRetrieve)
}

func VocabulariesRetrieve(c *gin.Context) {
	vocabs, err := getVocabularies()
	fmt.Println(vocabs)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("id", errors.New("Invalid ID"))) // need to change later
		return
	}
	// var vocabularies []VocabularyResponse
	// err = mapstructure.Decode(vocabs, &vocabularies)
	// if err != nil {
	// 	fmt.Println("Error mapping data:", err)
	// } 
	c.JSON(http.StatusOK, gin.H{"vocabs": vocabs})
}