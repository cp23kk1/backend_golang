package users

import (
	"cp23kk1/common/databases"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID          uint    `gorm:"primary_key"`
	Username    string  `gorm:"column:Username"`
	Email       string  `gorm:"column:email;unique_index"`
}

func AutoMigrate() {
	db := databases.Init()

	db.AutoMigrate(&UserModel{})
} 

func (UserModel) TableName() string {
    return "user"
}

//create a User
func create(data interface{}) (error) {
	db := databases.GetDB()
	err := db.Create(data).Error
	return err
}

//get User by id
func selectById(id int) (UserModel, error) {
	db := databases.GetDB()
	var user UserModel
	// err := db.Where("id = ?", id).First(&user).Error
	err := db.Raw("SELECT * FROM user WHERE id = ?", id).Scan(&user).Error
	return user, err
}

//update User
func (UserModel *UserModel) update(data interface{}) (error) {
	db := databases.GetDB()
	err := db.Model(UserModel).Save(data).Error
	return err
}

//delete User
func delete(condition interface{}) (UserModel, error) {
	db := databases.GetDB()
	var user UserModel
	err := db.Where(condition).Delete(&user).Error
	return user, err
}