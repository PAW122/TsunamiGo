package slashcommands

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// HandleAvatarCommand obsługuje slash command /avatar
func HandleAvatarCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Pobranie użytkownika z opcji slash command
	for _, option := range i.ApplicationCommandData().Options {
		if option.Type == discordgo.ApplicationCommandOptionUser {
			user, err := s.User(option.Value.(string))
			if err != nil {
				// Obsługa błędu
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Error: Unable to find user.",
					},
				})
				return
			}

			// Tworzenie embeda z awatarem użytkownika
			embed := &discordgo.MessageEmbed{
				Color:       0x0000ff, // Niebieski kolor
				Title:       "Avatar",
				Description: "Avatar użytkownika " + user.Username + "\n[Click here](" + user.AvatarURL("2048") + ") to see the avatar in higher resolution",
				Image: &discordgo.MessageEmbedImage{
					URL: user.AvatarURL("512"),
				},
				Timestamp: time.Now().Format(time.RFC3339),
			}

			// Wysłanie odpowiedzi na slash command z embedem
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
				},
			})

			return
		}
	}

	// Jeśli nie znaleziono opcji użytkownika
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Error: User not specified.",
		},
	})
}
