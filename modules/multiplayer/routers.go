package multiplayer

import (
	"cp23kk1/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMultiplayerRoutes(rg *gin.RouterGroup) {
	multiplayerRouter := rg.Group("/multiplayer")

	// gameplay.Use(auth.AuthMiddleware(true, "access_token"))
	multiplayerRouter.GET("create-lobby", CreateLobby)
	multiplayerRouter.GET("join-lobby", JoinLobby)
	multiplayerRouter.GET("get-lobby", GetLobby)
	multiplayerRouter.POST("update-lobby", UpdateLobby)
}

func JoinLobby(c *gin.Context) {
	ServeWs(c.Writer, c.Request, false)

}
func CreateLobby(c *gin.Context) {
	ServeWs(c.Writer, c.Request, true)
}
func GetLobby(c *gin.Context) {

	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"lobby": GetlobbyService()}))

}
func UpdateLobby(c *gin.Context) {
	updateModel := NewUpdateLobbyModelValidator()
	if err := updateModel.Bind(c); err != nil {

		c.JSON(http.StatusBadRequest, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "error"}, map[string]interface{}{"errorMessage": err.Error()}))
		return
	}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"lobby": UpdateRoomService(updateModel.RoomID, updateModel.IsPlayed)}))

}
