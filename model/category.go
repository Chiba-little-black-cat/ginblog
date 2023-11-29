package model

import (
	"errors"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key; auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCategoryName(name string) (bool, error) {
	var id int
	err := db.Model(&Category{}).Select("id").Where("name = ?", name).First(&id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Model(&Category{}).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
		return false, err
	}

	return true, nil
}

func CreateCategory(data *Category) error {
	err := db.Create(data).Error
	return err
}

func GetCategories(pageSize int, pageNum int) ([]Category, error) {
	var categories []Category
	err := db.Select("id, name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories).Error
	return categories, err
}

func EditCategory(id int, category *Category) error {
	err := db.Model(&Category{}).Where("id = ?", id).Update("name", category.Name).Error

	return err
}

func DeleteCategory(id int) error {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	return err
}
