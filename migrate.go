package seaottermsdb

import (
	"seaotterms-db/auth"
	"seaotterms-db/blog"
	"seaotterms-db/discordbot"
	"seaotterms-db/teach"

	"github.com/sirupsen/logrus"
)

func Migration(dbm *DBModel) {
	switch dbm.modelType {
	case BlogModel:
		dbm.DB.AutoMigrate(&blog.User{})
		dbm.DB.AutoMigrate(&blog.Tag{})
		dbm.DB.AutoMigrate(&blog.Article{})
		dbm.DB.AutoMigrate(&blog.Todo{})
		dbm.DB.AutoMigrate(&blog.TodoTopic{})
		dbm.DB.AutoMigrate(&blog.SystemTodo{})
	case DiscordBotModel:
		dbm.DB.AutoMigrate(&discordbot.Member{})
		dbm.DB.AutoMigrate(&discordbot.Log{})
		dbm.DB.AutoMigrate(&discordbot.DedicatedChannel{})
	case AuthModel:
		dbm.DB.AutoMigrate(&auth.Token{})
	case TeachModel:
		dbm.DB.AutoMigrate(&teach.Series{})
		dbm.DB.AutoMigrate(&teach.Article{})
		dbm.DB.AutoMigrate(&teach.Comment{})
	default:
		logrus.Warn("seaotterms-db: no matching migration found, skipping operation")
	}
}
