package user

import (
	"cp23kk1/common/databases"
	"cp23kk1/modules/repository/enum"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {

	db.AutoMigrate(&UserModel{})
}

func (UserModel) TableName() string {
	return "user"
}

func CreateUser(email string, role enum.Role, displayName string, isActive bool, image string, isPrivateProfile bool) {
	db := databases.GetDB()

	user := UserModel{
		Email:            &email,
		Role:             role,
		DisplayName:      &displayName,
		IsActive:         isActive,
		Image:            &image,
		IsPrivateProfile: isPrivateProfile,
	}
	db.Create(&user)
}

func FindUserByID(id int) (*UserModel, error) {
	db := databases.GetDB()
	var user UserModel
	err := db.Where("id = ?", id).Preload("ScoreBoards").First(&user).Error
	return &user, err
}

func UpdateUser(id int, email string, role enum.Role, displayName string, isActive bool, image string, isPrivateProfile bool) error {
	db := databases.GetDB()

	user, err := FindUserByID(id)
	if err != nil {
		return err
	}

	user.Email = &email
	user.Role = role
	user.DisplayName = &displayName
	user.IsActive = isActive
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
