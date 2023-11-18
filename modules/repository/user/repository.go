package user

import (
	"cp23kk1/common/databases"
	"cp23kk1/modules/repository/enum"
)

func CreateUser(email string, role enum.Role, displayName string, image string, isPrivateProfile bool) error {
	db := databases.GetDB()

	user := UserModel{
		Email:            &email,
		Role:             role,
		DisplayName:      &displayName,
		IsActive:         true,
		Image:            &image,
		IsPrivateProfile: isPrivateProfile,
	}
	return db.Create(&user).Error
}

func FindUserByID(id int) (*UserModel, error) {
	db := databases.GetDB()
	var user UserModel
	err := db.Where("id = ?", id).Preload("ScoreBoards").First(&user).Error
	return &user, err
}
func FindAllUsers() (*[]UserModel, error) {
	db := databases.GetDB()
	var user []UserModel
	err := db.Preload("ScoreBoards").Find(&user).Error
	return &user, err
}

func UpdateUser(id int, email string, role enum.Role, displayName string, image string, isPrivateProfile bool) error {
	db := databases.GetDB()

	user, err := FindUserByID(id)
	if err != nil {
		return err
	}

	user.Email = &email
	user.Role = role
	user.DisplayName = &displayName
	user.Image = &image
	user.IsPrivateProfile = isPrivateProfile

	if result := db.Save(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(id int) error {
	db := databases.GetDB()

	user, err := FindUserByID(id)
	if err != nil {
		return err
	}

	if result := db.Delete(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
