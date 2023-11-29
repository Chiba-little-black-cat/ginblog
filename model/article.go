package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title        string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int      `gorm:"type:int;not null" json:"cid"`
	Category     Category `gorm:"foreignKey:Cid"`
	Description  string   `gorm:"type:varchar(200)" json:"description"`
	Content      string   `gorm:"type:longtext" json:"content"`
	Img          string   `gorm:"type:varchar(100)" json:"img"`
	CommentCount int      `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int      `gorm:"type:int;not null;default:0" json:"read_count"`
}

func CreateArticle(article *Article) error {
	err := db.Create(&article).Error
	return err
}

// GetArticlesByCategoryId returns a list of articles with the given category id.
// It will not return the article contents !!
func GetArticlesByCategoryId(id int, pageSize int, pageNum int) ([]Article, error) {
	var articles []Article

	err := db.Preload("Category").
		Limit(pageSize).Offset((pageNum-1)*pageSize).
		Where("cid = ?", id).
		Find(&articles).Error

	return articles, err
}

func GetArticleCountByCategoryId(id int) (int64, error) {
	var count int64
	err := db.Preload("Category").Where("cid = ?", id).Count(&count).Error

	return count, err
}

func GetArticleById(id int) (Article, error) {
	var article Article

	err := db.Where("id = ?", id).
		Preload("Category").
		First(&article).Error

	db.Model(&article).
		Where("id = ?", id).
		UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))

	return article, err
}

// GetArticles returns a list of articles. It will not return the article contents !!
func GetArticles(pageSize int, pageNum int) ([]Article, error) {
	var articles []Article

	err := db.Select("articles.id, title, img, created_at, updated_at, description, comment_count, read_count, name").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("Created_At DESC").
		Joins("Category").
		Find(&articles).Error

	return articles, err
}

func GetArticleCount() (int64, error) {
	var count int64
	err := db.Model(&Article{}).Count(&count).Error
	return count, err
}

// SearchArticlesByTitle returns a list of articles with the given title.
// It will not return the article contents !!
func SearchArticlesByTitle(title string, pageSize int, pageNum int) ([]Article, error) {
	var articles []Article

	err := db.Select("articles.id,title, img, created_at, updated_at, description, comment_count, read_count, name").
		Order("Created_At DESC").
		Joins("Category").
		Where("title LIKE ?", title+"%").Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&articles).Error

	return articles, err
}

func GetArticleCountByTitle(title string) (int64, error) {
	var count int64
	err := db.Model(&Article{}).Where("title LIKE ?", title+"%").Count(&count).Error
	return count, err
}

func EditArticle(id int, article *Article) error {
	var maps = make(map[string]interface{})

	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["description"] = article.Description
	maps["content"] = article.Content
	maps["img"] = article.Img

	err := db.Model(&Article{}).Where("id = ? ", id).Updates(&maps).Error

	return err
}

func DeleteArticle(id int) error {
	err := db.Where("id = ? ", id).Delete(&Article{}).Error
	return err
}
