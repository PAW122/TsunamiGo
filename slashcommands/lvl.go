package slashcommands

import (
	"fmt"
	database "tsunamiBot/db"
	types "tsunamiBot/types"

	"github.com/bwmarrin/discordgo"
)

// HandlePingCommand obsługuje slash command /ping
func HandleLvlCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Pobierz identyfikator użytkownika i serwera z interakcji
	var userId, guildId string
	if i.Member != nil { // Sprawdź, czy Member jest zdefiniowany
		userId = i.Member.User.ID
		guildId = i.GuildID
	} else { // Jeśli Member nie jest zdefiniowany, korzystaj z autora interakcji
		userId = i.User.ID
		guildId = i.GuildID
	}
	// discordgo.ApplicationCommandInteractionDataOption

	// Sprawdź, czy opcja użytkownika została przekazana
	options := i.ApplicationCommandData().Options
	var targetUserId string
	for _, option := range options {
		if option.Name == "user" {
			targetUserId = option.Value.(string)
			break
		}
	}

	// Jeśli opcja użytkownika została przekazana, użyj tej wartości jako ID użytkownika
	if targetUserId != "" {
		userId = targetUserId
	}

	// Pobierz dane użytkownika
	userData := Get_user_lvl(userId, guildId)

	// Utwórz wiadomość zawierającą informacje o poziomie użytkownika
	User_lvl := userData.Lvl
	User_messages := userData.Messages
	Res := fmt.Sprintf("User lvl %d sended messages: %d", User_lvl, User_messages)

	// Wysyłanie odpowiedzi na slash command
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: Res,
		},
	})
}

func Get_user_lvl(userId string, guildId string) types.GetUserXpRes {
	req := types.GetUserXp{
		UserId:  userId,
		GuildId: guildId,
	}

	res := database.GetUserXp(req)
	return res
}
