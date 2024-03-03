package auth

import (
	"cp23kk1/common"
	"cp23kk1/common/config"
	"cp23kk1/common/databases"
	"cp23kk1/modules/repository/user"
	"time"

	"github.com/gin-gonic/gin"
)

func GuestLoginService() (*string, *string, error) {
	userRepository := user.NewUserRepository(databases.GetDB())

	displayName := "Guest"
	guestImage := "https://icons.veryicon.com/png/o/miscellaneous/youyinzhibo/guest.png"
	guestUser := &databases.UserModel{

		Email:       nil,
		DisplayName: &displayName,
		Image:       &guestImage,
		Role:        "user",
	}

	newUser, err := userRepository.CreateUser(*guestUser)
	config, _ := config.LoadConfig()

	// Generate Tokens
	access_token, err := common.CreateToken(config.AccessTokenExpiresIn, newUser.ID, config.AccessTokenPrivateKey)
	refresh_token, err := common.CreateToken(time.Duration(time.Hour*525600), newUser.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		return nil, nil, err
	}
	return &access_token, &refresh_token, nil

}

func GoogleOAuthService(c *gin.Context) (access, refresh string, err error) {
	userRepository := user.NewUserRepository(databases.GetDB())
	code := c.Query("code")
	if code == "" {

		return "", "", err
	}

	// Use the code to get the id and access tokens
	tokenRes, err := common.GetGoogleOauthToken(code)

	if err != nil {
		return "", "", err
	}

	userFromGoogle, err := common.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		return "", "", err
	}

	resBody := &databases.UserModel{

		Email:       &userFromGoogle.Email,
		DisplayName: &userFromGoogle.Name,
		Image:       &userFromGoogle.Picture,
		// Provider:    "google",
		Role: "user",
	}

	updatedUser, err := userRepository.Upsert(*resBody)
	if err != nil {
		return "", "", err
	}

	config, _ := config.LoadConfig()

	// Generate Tokens
	access_token, err := common.CreateToken(config.AccessTokenExpiresIn, updatedUser.ID, config.AccessTokenPrivateKey)
	if err != nil {
		return "", "", err
	}

	refresh_token, err := common.CreateToken(config.RefreshTokenExpiresIn, updatedUser.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		return "", "", err
	}
	return access_token, refresh_token, err
}
