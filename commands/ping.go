package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Define the ping command
var PingCommand = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Replies with Pong!",
}

// Handle the ping command interaction
func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
	if err != nil {
		fmt.Printf("Cannot respond to interaction: %v\n", err)
	}
}
