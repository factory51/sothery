package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// List of all commands
var AllCommands = []*Command{
	{PingCommand, PingHandler},
	// Add other commands here
}

func RegisterCommands(s *discordgo.Session, guildID string) {
	for _, cmd := range AllCommands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, cmd.Definition)
		if err != nil {
			fmt.Printf("Cannot create '%v' command: %v\n", cmd.Definition.Name, err)
		}
	}
}

func UnregisterCommands(s *discordgo.Session, guildID string) {
	commands, err := s.ApplicationCommands(s.State.User.ID, guildID)
	if err != nil {
		fmt.Printf("Cannot fetch commands: %v\n", err)
		return
	}

	for _, cmd := range commands {
		err := s.ApplicationCommandDelete(s.State.User.ID, guildID, cmd.ID)
		if err != nil {
			fmt.Printf("Cannot delete '%v' command: %v\n", cmd.Name, err)
		}
	}
}
