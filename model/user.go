package model

import (
	"errors"
	"ginblog/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(60);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

func CheckUsername(username string) (bool, error) {
	var id int
	err := db.Model(&User{}).Select("id").Where("username = ?", username).First(&id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Model(&User{}).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
		return false, err
	}

	return true, nil
}

func (user *User) BeforeCreate(_ *gorm.DB) error {
	// Bcrypt password
	user.Password, _ = utils.BcryptPassWord(user.Password)
	user.Role = 2
	return nil
}

func (user *User) BeforeUpdate(_ *gorm.DB) (err error) {
	// Bcrypt password
	user.Password, _ = utils.BcryptPassWord(user.Password)
	return nil
}

func CreateUser(data *User) error {
	err := db.Create(data).Error
	return err
}

func GetUsers(pageSize int, pageNum int) ([]User, error) {
	var users []User
	err := db.Select("id, username, role, created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	return users, err
}

func EditUser(id int, user *User) error {
	var maps = make(map[string]interface{})
	maps["username"] = user.Username
	maps["role"] = user.Role

	err := db.Model(&User{}).Where("id = ?", id).Updates(maps).Error

	return err
}

func DeleteUser(id int) error {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	return err
}

func GetPasswordByUsername(userName string) (string, error) {
	var password string
	err = db.Model(&User{}).Select("username").Where("username = ?", userName).Find(&password).Error
	return password, err
}
