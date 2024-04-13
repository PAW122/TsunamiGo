package slashcommands

import (
	"github.com/bwmarrin/discordgo"
)

// HelpCommandHandler obsługuje slash command /help
func HelpCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Pobieranie mapy komend z plików
	commandMap := GetCommandMap()

	// Pobranie wybranej komendy (jeśli podana)
	chosenCommand := i.ApplicationCommandData().Options[0].StringValue()

	// Sprawdzenie, czy wybrana komenda jest dostępna
	if chosenCommand != "" {
		// Wyświetlenie opisu wybranej komendy
		if helpMessage, ok := commandMap[chosenCommand]; ok {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: helpMessage,
				},
			})
		} else {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Nie znaleziono komendy o podanej nazwie.",
				},
			})
		}
	} else {
		// Wyświetlenie listy wszystkich dostępnych komend
		var content string
		for command, helpMessage := range commandMap {
			content += "**/" + command + "**: " + helpMessage + "\n"
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	}
}

// GetCommandMap zwraca mapę komend i ich opisów
func GetCommandMap() map[string]string {
	// Tutaj możesz zaimplementować logikę wczytywania mapy komend z plików
	// Na potrzeby tego przykładu, zwrócę statyczną mapę
	commandMap := map[string]string{
		"ping":   "Pong! Odpowiada na komendę ping.",
		"avatar": "Wyświetla awatar wybranego użytkownika.",
		"8ball":  "Wysyła przypadkową odpowiedz",
		// Dodaj więcej komend wraz z ich opisami
	}
	return commandMap
}
