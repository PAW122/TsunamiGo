package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	// commands "tsunamiBot/commands"
	handler "tsunamiBot/handlers"

	"github.com/bwmarrin/discordgo"
)

var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "Nzk3MDcwODA2ODg1OTkwNDMx.GNeuIa.iLqAvMdiXzPiK8X2lHqOANt2a8iTnrRiY5Tpe4", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var client *discordgo.Session

func init() { flag.Parse() }

func main() {
	// Tworzenie nowego połączenia Discord
	client, err := discordgo.New("Bot " + *BotToken)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Rejestrowanie funkcji obsługi komend
	client.AddHandler(handler.CommandHandler)
	client.AddHandler(interactionCreate)
	client.AddHandler(handler.Lvl_handler)

	// Otwieranie połączenia
	err = client.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	// Oczekiwanie na sygnał zakończenia
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Zamykanie połączenia
	client.Close()
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Sprawdzanie czy to jest INTERACTION_CREATE dla slash command
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	// Obsługa slash command
	if handler, ok := handler.SlashCommandHandler[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
	// Dodaj więcej obsługi dla innych slash command
}
