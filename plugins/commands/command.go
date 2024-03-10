package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shaunsational/go-bot/lib/session"
)

/*
//	Description : A Basic command "module" to add to the bot
*/

func Cmdname(s *discordgo.Session, conf *session.Config) ([]*discordgo.ApplicationCommand) {
//	 ———▲——— This function will need to be added to the init() of lib/commands/commands.go
	var cmdName = "command-name" // this is used on lines 18 & 25

	pkgCommands := []*discordgo.ApplicationCommand{
		{
			Name:			cmdName,
			Description:	"this is a basic slash command",
		},
	}

	s.AddHandler(func (s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			if i.ApplicationCommandData().Name == cmdName {

				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData {
						Content:	"...",

						// the following line makes it an ephemeral message,
						// it is only shown to the user exeucting the command.
						// it is not required.
						Flags:		discordgo.MessageFlagsEphemeral,
					},
				})

				if err != nil {
					// Handle the error
				}
			}
		}
	})

	return pkgCommands
}