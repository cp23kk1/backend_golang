package users

import (
	"cp23kk1/common/databases"
	userRepo "cp23kk1/modules/repository/user"
)

type UserUseCase struct{}

func getUser(id int) (*userRepo.UserModel, error) {
	userRepository := userRepo.NewUserRepository(databases.GetDB())
	return userRepository.FindUserByID(id)
}
