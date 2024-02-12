package gameplays

import (
	"net/http"

	"cp23kk1/common"
	"cp23kk1/modules/auth"

	"github.com/gin-gonic/gin"
)

func AddGameplayRoutes(rg *gin.RouterGroup) {
	gameplay := rg.Group("/gameplays")

	gameplay.Use(auth.AuthMiddleware(true, "access_token"))
	gameplay.GET("/vocabulary", RandomVocabularyForGamePlay)
}

func VocabulariesRetrieve(c *gin.Context) {
	vocabs, err := getVocabularies()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("eiei", err)) // need to change later
		return
	}
	c.JSON(http.StatusOK, gin.H{"vocabs": vocabs})
}

func RandomVocabularyForGamePlay(c *gin.Context) {
	vocabs, err := randomFromGamePlay()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("eiei", err)) // need to change later
		return
	}
	serializer := VocabsSerealizer{c, vocabs}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"vocabs": serializer.Response()}))
}
