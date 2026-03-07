package seaottermsdb

import (
	"seaotterms-db/auth"
	"seaotterms-db/blog"
	"seaotterms-db/discordbot"
	"seaotterms-db/teach"
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
		dbm.DB.AutoMigrate(&auth.User{})
	case TeachModel:
		dbm.DB.AutoMigrate(&teach.Series{})
		dbm.DB.AutoMigrate(&teach.Article{})
		dbm.DB.AutoMigrate(&teach.Comment{})
	}
}
