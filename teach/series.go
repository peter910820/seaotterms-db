package teach

import (
	"time"

	"gorm.io/gorm"
)

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
