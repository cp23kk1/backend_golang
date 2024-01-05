package auth

import (
	"cp23kk1/common"
	"cp23kk1/common/config"
	"cp23kk1/common/databases"
	"cp23kk1/modules/repository/user"
)

func GuestLoginService() (*string, *string, error) {
	userRepository := user.NewUserRepository(databases.GetDB())

	displayName := "Guest"
	guestUser := &user.UserModel{

		Email:       nil,
		DisplayName: &displayName,
		Image:       nil,
		Role:        "guest",
	}

	newUser, err := userRepository.CreateUser(*guestUser)
	config, _ := config.LoadConfig()

	// Generate Tokens
	access_token, err := common.CreateToken(config.AccessTokenExpiresIn, newUser.ID, config.AccessTokenPrivateKey)
	refresh_token, err := common.CreateToken(0, newUser.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		return nil, nil, err
	}
	return &access_token, &refresh_token, nil

}
