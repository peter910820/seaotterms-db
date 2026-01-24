package seaottermsdb

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ModelSet string

type (
	ConnectDBConfig struct {
		Owner    string
		Password string
		DBName   string
		Port     int
	}
	DBModel struct {
		DB        *gorm.DB
		modelType ModelSet
	}
)

const (
	BlogModel       ModelSet = "Blog"
	DiscordBotModel ModelSet = "DiscordBot"
	AuthModel       ModelSet = "Auth"
	TeachModel      ModelSet = "Teach"
)

func InitDsn(config ConnectDBConfig) (*DBModel, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Owner,
		config.Password,
		config.DBName,
		config.Port)

	// get connect db variable
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// 連接資料庫失敗
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		// 無法取得sql.DB
		return nil, err
	}

	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	dbModel := DBModel{
		DB: db,
	}
	switch {
	case strings.HasSuffix(config.DBName, "Blog"):
		dbModel.modelType = BlogModel
	case strings.HasSuffix(config.DBName, "DiscordBot"):
		dbModel.modelType = DiscordBotModel
	case strings.HasSuffix(config.DBName, "Auth"):
		dbModel.modelType = AuthModel
	case strings.HasSuffix(config.DBName, "Teach"):
		dbModel.modelType = TeachModel
	default:
		dbModel.modelType = ""
	}

	return &dbModel, nil
}
