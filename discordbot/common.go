package discordbot

type ChannelType string

const (
	LevelupChannel     ChannelType = "LEVELUP" // level up channel
	BotChannel         ChannelType = "BOT"
	VoiceManageChannel ChannelType = "VOICE_MANAGE"
)
