package routes

import (
	"cp23kk1/modules/cms/passage"
	"cp23kk1/modules/cms/passage_history"
	"cp23kk1/modules/cms/score_board"
	"cp23kk1/modules/cms/sentence"
	"cp23kk1/modules/cms/vocabulary"
	"cp23kk1/modules/cms/vocabulary_history"
	"cp23kk1/modules/gameplays"
	"cp23kk1/modules/history"
	"cp23kk1/modules/ping"
	"cp23kk1/modules/users"

	"github.com/gin-gonic/gin"
)

// Run will start the server
func Run(router *gin.Engine) {
	getRoutes(router)
	router.Run(":8080")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes(router *gin.Engine) {
	api := router.Group("/api")
	ping.AddPingRoutes(api)
	users.AddUserRoutes(api)
	gameplays.AddGameplayRoutes(api)
	history.AddHistoryRoutes(api)

	v1 := router.Group("/api/cms")
	passage.SetupPassageRoutes(v1)
	passage_history.SetupPassageHistoryRoutes(v1)
	vocabulary.SetupVocabularyRoutes(v1)
	vocabulary_history.SetupVocabularyHistoryRoutes(v1)
	score_board.SetupScoreBoardRoutes(v1)
	sentence.SetupSentenceRoutes(v1)

}
