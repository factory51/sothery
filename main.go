package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"factory51/sotherby/src/commands" // Adjust the import path as necessary

	"github.com/bwmarrin/discordgo"
)

var (
	Token   = ""
	GuildID = "" // Set to "" if you want global commands
)

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(interactionCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Register slash commands
	commands.RegisterCommands(dg, GuildID)

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Unregister slash commands
	commands.UnregisterCommands(dg, GuildID)

	dg.Close()
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	for _, cmd := range commands.AllCommands {
		if i.ApplicationCommandData().Name == cmd.Definition.Name {
			cmd.Handler(s, i)
			return
		}
	}
}
