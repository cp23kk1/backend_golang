package auth

import (
	"cp23kk1/common"
	"cp23kk1/common/config"
	"cp23kk1/common/databases"
	"cp23kk1/modules/repository/user"
	"fmt"
	"net/http"

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
	return
}

func RefreshToken(c *gin.Context) {
	config, _ := config.LoadConfig()
	userId, exists := c.Get("userId")
	fmt.Println("userId:  ", userId)
	if exists == false {
		c.JSON(http.StatusUnauthorized, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Status: "failed", Message: "Error when retrived token"}, map[string]interface{}{}))
	}
	access_token, err := common.CreateToken(config.AccessTokenExpiresIn, userId, config.AccessTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60*60, "/", config.ORIGIN, false, true)

}
func GoogleOAuth(c *gin.Context) {
	userRepository := user.NewUserRepository(databases.GetDB())

	code := c.Query("code")
	var pathUrl string = "/"

	if c.Query("state") != "" {
		pathUrl = c.Query("state")
	}

	if code == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}

	// Use the code to get the id and access tokens
	tokenRes, err := common.GetGoogleOauthToken(code)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	userFromGoogle, err := common.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	resBody := &user.UserModel{

		Email:       &userFromGoogle.Email,
		DisplayName: &userFromGoogle.Name,
		Image:       &userFromGoogle.Picture,
		// Provider:    "google",
		Role: "user",
	}

	updatedUser, err := userRepository.Upsert(*resBody)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}

	config, _ := config.LoadConfig()

	// Generate Tokens
	access_token, err := common.CreateToken(config.AccessTokenExpiresIn, updatedUser.ID, config.AccessTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := common.CreateToken(config.RefreshTokenExpiresIn, updatedUser.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60*60, "/", config.ORIGIN, false, true)
	c.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60*60, "/", config.ORIGIN, false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60*60, "/", config.ORIGIN, false, false)
	var basePath = ""
	if config.ENV != "prod" {
		basePath = config.ENV
	}
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