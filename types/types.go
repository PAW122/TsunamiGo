package types

import "github.com/bwmarrin/discordgo"

// Command represents a command
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
}

type AddXp struct {
	GuildId     string `json:"guildId"`
	UserId      string `json:"userId"`
	AddXp       int    `json:"addXp"`
	NewMessages int    `json:"newMessages"`
}

type GetUserXp struct {
	UserId  string `json:"userId"`
	GuildId string `json:"guildId"`
}

type GetUserXpRes struct {
	UserId   string `json:"userId"`
	GuildId  string `json:"guildId"`
	Xp       int    `json:"xp"`
	Lvl      int    `json:"lvl"`
	Messages int    `json:"messages"`
	Error    error  `json:"error"`
}
