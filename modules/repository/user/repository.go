package user

import (
	"cp23kk1/common/databases"
	"cp23kk1/common/enum"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (u UserRepository) CreateUser(user databases.UserModel) (databases.UserModel, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return databases.UserModel{}, err
	}
	return user, nil
}

func (u UserRepository) FindUserByID(id uint) (*databases.UserModel, error) {
	var user databases.UserModel
	err := u.db.Where("id = ?", id).Preload("ScoreBoards").First(&user).Error
	return &user, err
}

func (u UserRepository) FindUserByEmail(email string) (*databases.UserModel, error) {
	var user databases.UserModel
	err := u.db.Where("email = ?", email).Preload("ScoreBoards").First(&user).Error
	return &user, err
}
func (u UserRepository) FindAllUsers() (*[]databases.UserModel, error) {
	var user []databases.UserModel
	err := u.db.Preload("ScoreBoards").Find(&user).Error
	return &user, err
}

func (u UserRepository) UpdateUser(id uint, email *string, role enum.Role, displayName string, image *string, isPrivateProfile bool) (*databases.UserModel, error) {

	user, err := u.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	user.Email = email
	user.Role = role
	user.DisplayName = &displayName
	user.Image = image
	user.IsPrivateProfile = isPrivateProfile

	if result := u.db.Save(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (u UserRepository) UpdateDislayNameUser(id uint, displayName string) error {

	user, err := u.FindUserByID(id)
	if err != nil {
		return err
	}

	user.DisplayName = &displayName

	if result := u.db.Save(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u UserRepository) DeleteUser(id uint) error {

	user, err := u.FindUserByID(id)
	if err != nil {
		return err
	}

	if result := u.db.Delete(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u UserRepository) Upsert(newUser databases.UserModel) (databases.UserModel, error) {

	var existingUser databases.UserModel
	result := u.db.Where(databases.UserModel{Email: newUser.Email}).First(&existingUser)

	// If the user doesn't exist, create a new one
	if result.Error != nil {
		fmt.Print("NOTFOUNDNDDD")
		if result.Error == gorm.ErrRecordNotFound {
			u.db.Create(&newUser)
			fmt.Println("User created:", newUser)
			return newUser, nil
		} else {
			return databases.UserModel{}, result.Error
		}
	} else {
		// If the user already exists, update the existing record
		u.db.Model(&existingUser).Updates(newUser)
		fmt.Println("User updated:", existingUser)
		return existingUser, nil
	}
}
