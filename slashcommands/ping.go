package slashcommands

import (
	"github.com/bwmarrin/discordgo"
)

// HandlePingCommand obsługuje slash command /ping
func HandlePingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Wysyłanie odpowiedzi na slash command
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,

		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
