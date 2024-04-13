package types

import "github.com/bwmarrin/discordgo"

// Command represents a command
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
}
