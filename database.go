package seaottermsdb

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"seaotterms-db/auth"
	"seaotterms-db/discordbot"
)

type ConnectDBConfig struct {
	Owner    string
	Password string
	DBName   string
	Port     int
}

type DBModel struct {
	DB        *gorm.DB
	modelsSet string
}

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
	case strings.HasSuffix(config.DBName, "DiscordBot"):
		dbModel.modelsSet = "DiscordBot"
	case strings.HasSuffix(config.DBName, "Auth"):
		dbModel.modelsSet = "Auth"
	default:
		dbModel.modelsSet = ""
	}

	return &dbModel, nil
}

func Migration(dbm *DBModel) {
	switch dbm.modelsSet {
	case "Auth":
		dbm.DB.AutoMigrate(&auth.Token{})
	case "DiscordBot":
		dbm.DB.AutoMigrate(&discordbot.Member{})
		dbm.DB.AutoMigrate(&discordbot.Log{})
		dbm.DB.AutoMigrate(&discordbot.DedicatedChannel{})
	}
}
