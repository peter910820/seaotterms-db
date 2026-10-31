package teach

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Article struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"NOT NULL" json:"title"`
	Image       string         `json:"image"`
	SeriesID    uint           `gorm:"NOT NULL" json:"seriesId"` // series table id
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags"`
	Content     string         `gorm:"NOT NULL" json:"content"`
	CreatedAt   time.Time      `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	CreatedName string         `gorm:"NOT NULL" json:"createdName"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	UpdatedName string         `json:"updatedName"`
}

// 用ID查詢Article
//
// 找不到資料回傳錯誤
func FindArticleByID(dbs *gorm.DB, id string) (*Article, error) {
	var data Article
	err := dbs.Scopes(latestOrder).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// 用seriesID查詢Article
func FindArticleBySeriesID(dbs *gorm.DB, seriesID string) ([]Article, error) {
	var data []Article

	err := dbs.Scopes(latestOrder).Where("series_id = ?", seriesID).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 取得全部的Article
func FindAllArticle(dbs *gorm.DB) ([]Article, error) {
	var data []Article

	err := dbs.Scopes(latestOrder).Find(&data).Error
	return data, err
}

// 創建新的Article
func CreateArticle(dbs *gorm.DB, data *Article) error {
	return dbs.Create(&data).Error
}

// 更新指定的Article
func UpdateArticle(dbs *gorm.DB, id string, data *Article) error {
	return dbs.Model(&Article{}).Where("id = ?", id).
		Select("title", "image", "series_id", "tags", "content", "updated_at", "updated_name").
		Updates(data).Error
}
