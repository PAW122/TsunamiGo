package commands

import "github.com/bwmarrin/discordgo"

// PingCommand obsługuje komendę !ping
type PingCommand struct{}

// Execute obsługuje wykonanie komendy
func (c *PingCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}
