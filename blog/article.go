package blog

import (
	"time"

	"gorm.io/gorm"
)

// all article
type Article struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"NOT NULL" json:"title"`
	Content   string    `gorm:"NOT NULL" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tags      []Tag     `gorm:"many2many:article_tags" json:"tags"`
}

// 查詢所有文章，並載入標籤
func FindAllArticles(db *gorm.DB) ([]Article, error) {
	var data []Article

	err := db.Preload("Tags").Order("created_at desc").Find(&data).Error
	return data, err
}

// 依標籤名稱查詢文章
func FindArticlesByTag(db *gorm.DB, name string) ([]Article, error) {
	var data []Article

	err := db.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Joins("JOIN tags ON tags.name = article_tags.tag_name").
		Where("tags.name = ?", name).Find(&data).Error
	return data, err
}

// 新增文章
func CreateArticle(db *gorm.DB, data *Article) error {
	return db.Create(data).Error
}

// 更新文章標題與內容
func UpdateArticle(db *gorm.DB, article *Article) error {
	return db.Model(article).
		Select("title", "content", "updated_at").
		Updates(article).Error
}

// 替換文章的標籤關聯
func ReplaceArticleTags(db *gorm.DB, article *Article, tags []Tag) error {
	return db.Model(article).Association("Tags").Replace(tags)
}

// 依 ID 查詢文章，並載入標籤
func FindArticleByID(db *gorm.DB, id string) (*Article, error) {
	var article Article

	if err := db.Preload("Tags").First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// 清除文章的標籤關聯
func ClearArticleTags(db *gorm.DB, article *Article) error {
	return db.Model(article).Association("Tags").Clear()
}

// 刪除文章
func DeleteArticle(db *gorm.DB, article *Article) error {
	return db.Delete(article).Error
}
