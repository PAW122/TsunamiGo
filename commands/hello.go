package commands

import (
	"github.com/bwmarrin/discordgo"
)

type HelloCommand struct{}

// Execute obsługuje wykonanie komendy
func (c *HelloCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Username+"!")
}
