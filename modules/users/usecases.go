package users

import userRepo "cp23kk1/modules/repository/user"

type UserUseCase struct{}

func getUser(id int) (*userRepo.UserModel, error) {
	return userRepo.FindUserByID(id)
}
