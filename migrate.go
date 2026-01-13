package seaottermsdb

import (
	"seaotterms-db/auth"
	"seaotterms-db/discordbot"

	"github.com/sirupsen/logrus"
)

func Migration(dbm *DBModel) {
	switch dbm.modelsSet {
	case "Auth":
		dbm.DB.AutoMigrate(&auth.Token{})
	case "DiscordBot":
		dbm.DB.AutoMigrate(&discordbot.Member{})
		dbm.DB.AutoMigrate(&discordbot.Log{})
		dbm.DB.AutoMigrate(&discordbot.DedicatedChannel{})
	default:
		logrus.Warn("seaotterms-db: no matching migration found, skipping operation")
	}
}
