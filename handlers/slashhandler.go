package handler

import (
	slashcommands "tsunamiBot/slashcommands"

	"github.com/bwmarrin/discordgo"
)

// SlashCommandHandler jest mapą komend dla slash commands
var SlashCommandHandler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping":   slashcommands.HandlePingCommand,
	"8ball":  slashcommands.HandleEightBallCommand,
	"avatar": slashcommands.HandleAvatarCommand,
	"help":   slashcommands.HelpCommandHandler,
	// Dodaj więcej komend tutaj
}
