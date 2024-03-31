package auth

import (
	"cp23kk1/common"
	"cp23kk1/common/config"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	auth.GET("/google", GoogleOAuth)
	auth.POST("/guest", GuestLogin)
	refreshAuthRoute := auth.Use(AuthMiddleware(true, "refresh_token"))
	refreshAuthRoute.POST("/refresh", RefreshToken)
	authRoute := auth.Use(AuthMiddleware(true, "access_token"))
	authRoute.POST("/sign-out", logout)

}
func logout(c *gin.Context) {

	c.SetCookie("access_token", "", 0, "/", "", true, true)
	c.SetCookie("refresh_token", "", 0, "/", "", true, true)
	c.SetCookie("logged_in", "", 0, "/", "", true, false)
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Status: "success", Message: "Logout Successfully"}, map[string]interface{}{}))
}

func RefreshToken(c *gin.Context) {
	config, _ := config.LoadConfig()
	userId, exists := c.Get("userId")
	fmt.Println("userId:  ", userId)
	if exists == false {
		c.JSON(http.StatusUnauthorized, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Status: "failed", Message: "Error when retrived token"}, map[string]interface{}{}))
		return
	}
	access_token, err := common.CreateToken(config.AccessTokenExpiresIn, userId, config.AccessTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60*60, "/", config.ORIGIN, false, true)
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Status: "success", Message: "Refresh token Successfully"}, map[string]interface{}{}))

}
func GoogleOAuth(c *gin.Context) {
	config, err := config.LoadConfig()
	var pathUrl string = "/"

	if c.Query("state") != "" {
		pathUrl = c.Query("state")
	}
	var basePath = ""
	if config.ENV != "prod" {
		basePath = config.ENV
	}
	urlA, err := url.Parse(config.ORIGIN + "/" + basePath + pathUrl)
	access_token, refresh_token, err := GoogleOAuthService(c)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	fmt.Println("userId: ", urlA.Query().Get("id"))
	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60*60, "/", config.ORIGIN, false, true)
	c.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60*60, "/", config.ORIGIN, false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60*60, "/", config.ORIGIN, false, false)

	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprint(config.ORIGIN+"/"+basePath, pathUrl))
}
func GuestLogin(c *gin.Context) {
	config, err := config.LoadConfig()
	access_token, refresh_token, err := GuestLoginService()
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Status: "failed", Message: err.Error()}, map[string]interface{}{}))
		return
	}
	c.SetCookie("access_token", *access_token, config.AccessTokenMaxAge*60*60, "/", config.ORIGIN, false, true)
	c.SetCookie("refresh_token", *refresh_token, 0, "/", config.ORIGIN, false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60*60, "/", config.ORIGIN, false, false)
	c.JSON(http.StatusCreated, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Status: "success", Message: "GuestUser created"}, map[string]interface{}{}))
}
