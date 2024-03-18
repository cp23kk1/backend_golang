package multiplayer

import (
	"github.com/gin-gonic/gin"
)

func AddMultiplayerRoutes(rg *gin.RouterGroup) {
	multiplayerRouter := rg.Group("/multiplayer")

	// gameplay.Use(auth.AuthMiddleware(true, "access_token"))
	multiplayerRouter.GET("create-lobby", CreateLobby)
	multiplayerRouter.GET("join-lobby", JoinLobby)
}

func JoinLobby(c *gin.Context) {
	ServeWs(c.Writer, c.Request, false)

}
func CreateLobby(c *gin.Context) {
	ServeWs(c.Writer, c.Request, true)

}
