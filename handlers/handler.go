package handler

import (
	"strings"
	commands "tsunamiBot/commands"
	types "tsunamiBot/types"

	"github.com/bwmarrin/discordgo"
)

func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	const prefix = ">"

	var commandsMap = map[string]types.Command{
		"hello": &commands.HelloCommand{},
		"ping":  &commands.PingCommand{},
		// Dodaj więcej komend tutaj
	}

	// Sprawdzanie czy wiadomość rozpoczyna się od prefiksu
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	// Podział treści wiadomości na komendę i argumenty
	args := strings.Fields(m.Content[len(prefix):])
	if len(args) == 0 {
		return
	}

	// Sprawdzanie, która komenda została wywołana
	if cmd, ok := commandsMap[args[0]]; ok {
		cmd.Execute(s, m)
	}
}
