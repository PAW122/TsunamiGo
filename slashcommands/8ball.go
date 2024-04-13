package slashcommands

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Lista fortun
var fortunes = []string{
	"**Yes**",
	"**No**",
	"**Maybe**",
	"**I don't know**",
	"**Probably**",
	"**I guess**",
	"**I'm not sure**",
	"**Surely**",
}

// HandleEightBallCommand obsługuje slash command /eightball
func HandleEightBallCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Losowanie jednej z fortun
	rand.Seed(time.Now().UnixNano())
	fortune := fortunes[rand.Intn(len(fortunes))]

	// Wysyłanie odpowiedzi na slash command
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fortune, // Przekazanie wylosowanej fortuny do Content
		},
	})
}
