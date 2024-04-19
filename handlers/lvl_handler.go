package handler

import (
	database "tsunamiBot/db"
	types "tsunamiBot/types"

	"github.com/bwmarrin/discordgo"
)

func Lvl_handler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author != nil && m.Author.Bot {
		return
	}

	data := types.AddXp{
		GuildId:     m.GuildID,
		UserId:      m.Author.ID,
		AddXp:       5,
		NewMessages: 1,
	}

	database.SaveAddXp(data)
}
