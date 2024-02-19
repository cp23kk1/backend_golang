package routes

import (
	"cp23kk1/common/config"
	"cp23kk1/modules/auth"
	"cp23kk1/modules/cms/passage"
	"cp23kk1/modules/cms/passage_history"
	"cp23kk1/modules/cms/score_board"
	"cp23kk1/modules/cms/sentence"
	"cp23kk1/modules/cms/sentence_history"
	"cp23kk1/modules/cms/user"
	"cp23kk1/modules/cms/vocabulary"
	"cp23kk1/modules/cms/vocabulary_history"
	"cp23kk1/modules/cms/vocabulary_related"
	"cp23kk1/modules/gameplays"
	"cp23kk1/modules/history"
	"cp23kk1/modules/ping"
	"cp23kk1/modules/score"
	"cp23kk1/modules/users"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Run will start the server
func Run(router *gin.Engine) {
	// config, _ := config.LoadConfig()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))
	getRoutes(router)
	router.Run(":8080")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes(router *gin.Engine) {
	config, _ := config.LoadConfig()
	api := router.Group("")
	if env := config.ENV; env == "prod" {
		api = router.Group("/api")
	} else {
		api = router.Group("/" + env + "/api")
	}

	ping.AddPingRoutes(api)
	users.AddUserRoutes(api)
	gameplays.AddGameplayRoutes(api)
	history.AddHistoryRoutes(api)
	auth.AddAuthRoutes(api)
	score.AddScoreRoutes(api)

	v1 := api.Group("/cms")
	passage.SetupPassageRoutes(v1)
	passage_history.SetupPassageHistoryRoutes(v1)
	vocabulary.SetupVocabularyRoutes(v1)
	vocabulary_history.SetupVocabularyHistoryRoutes(v1)
	score_board.SetupScoreBoardRoutes(v1)
	sentence.SetupSentenceRoutes(v1)
	sentence_history.SetupSentenceHistoryRoutes(v1)
	user.SetupUserRoutes(v1)
	vocabulary_related.SetupVocabularyRelatedRoutes(v1)

}
