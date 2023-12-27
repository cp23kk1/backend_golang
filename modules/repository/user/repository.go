package user

import (
	"cp23kk1/modules/repository/enum"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) CreateUser(email *string, role enum.Role, displayName string, image *string, isPrivateProfile bool) error {

	user := UserModel{
		Email:            email,
		Role:             role,
		DisplayName:      &displayName,
		IsActive:         true,
		Image:            image,
		IsPrivateProfile: isPrivateProfile,
	}
	return u.db.Create(&user).Error
}

func (u UserRepository) FindUserByID(id int) (*UserModel, error) {
	var user UserModel
	err := u.db.Where("id = ?", id).Preload("ScoreBoards").First(&user).Error
	return &user, err
}
func (u UserRepository) FindAllUsers() (*[]UserModel, error) {
	var user []UserModel
	err := u.db.Preload("ScoreBoards").Find(&user).Error
	return &user, err
}

func (u UserRepository) UpdateUser(id int, email *string, role enum.Role, displayName string, image *string, isPrivateProfile bool) error {

	user, err := u.FindUserByID(id)
	if err != nil {
		return err
	}

	user.Email = email
	user.Role = role
	user.DisplayName = &displayName
	user.Image = image
	user.IsPrivateProfile = isPrivateProfile

	if result := u.db.Save(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u UserRepository) DeleteUser(id int) error {

	user, err := u.FindUserByID(id)
	if err != nil {
		return err
	}

	if result := u.db.Delete(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u UserRepository) Upsert(newUser UserModel) (UserModel, error) {

	var existingUser UserModel
	result := u.db.Where(UserModel{Email: newUser.Email}).First(&existingUser)

	// If the user doesn't exist, create a new one
	if result.Error != nil {
		fmt.Print("NOTFOUNDNDDD")
		if result.Error == gorm.ErrRecordNotFound {
			u.db.Create(&newUser)
			fmt.Println("User created:", newUser)
			return newUser, nil
		} else {
			return UserModel{}, result.Error
		}
	} else {
		// If the user already exists, update the existing record
		u.db.Model(&existingUser).Updates(newUser)
		fmt.Println("User updated:", existingUser)
		return existingUser, nil
	}
}
