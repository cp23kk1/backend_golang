package routes

import (
	"cp23kk1/modules/ping"
	"cp23kk1/modules/users"
	"cp23kk1/modules/gameplays"

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
}