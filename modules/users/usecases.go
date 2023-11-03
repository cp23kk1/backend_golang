package users

type UserUseCase struct{}

func getUser(id int) (UserModel, error) {
	return selectById(id)
}
