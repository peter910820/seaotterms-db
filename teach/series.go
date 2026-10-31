package teach

import (
	"time"

	"gorm.io/gorm"
)

type Series struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Title         string    `gorm:"NOT NULL" json:"title"`
	Image         string    `json:"image"`
	ArticleAmount uint      `gorm:"NOT NULL; default:0" json:"articleAmount"`
	CreatedAt     time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	CreatedName   string    `gorm:"NOT NULL" json:"createdName"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedName   string    `json:"updatedName"`
}

func UpdateSeriesArticleAmount(db *gorm.DB, id uint, amount uint, updateTime time.Time) error {
	return db.Model(&Series{}).Where("id = ?", id).
		Select("article_amount", "updated_at").
		Updates(struct {
			ArticleAmount uint
			UpdateTime    time.Time
		}{amount, updateTime}).Error
}

// 用ID查詢Series
//
// 找不到資料回傳錯誤
func FindSeriesByID(dbs *gorm.DB, id string) (*Series, error) {
	var data Series

	err := dbs.Scopes(latestOrder).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// 取得全部的Series
func FindAllSeries(dbs *gorm.DB) ([]Series, error) {
	var data []Series

	err := dbs.Scopes(latestOrder).Find(&data).Error
	return data, err
}

// 創建新的Series
func CreateSeries(dbs *gorm.DB, data *Series) error {
	return dbs.Create(&data).Error
}

// 更新指定的Series
func UpdateSeries(dbs *gorm.DB, id string, data *Series) error {
	return dbs.Model(&Series{}).Where("id = ?", id).
		Select("title", "image", "updated_at", "updated_name").
		Updates(data).Error
}
