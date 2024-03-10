package commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/shaunsational/go-bot/plugins/commands"

	s "github.com/shaunsational/go-bot/lib/session"
)

var slash = []*discordgo.ApplicationCommand{}

func init() {
	// YOU WILL NEED TO ADD ALL COMMANDS HERE UNTIL I CAN FIND A WAY TO SELF-INSTANTIATE MODULES
	slash = append(slash, commands.Cmdname(s.Discord, s.Conf)...)

	s.Discord.AddHandler(messageCreate)
}

func Load() ([]*discordgo.ApplicationCommand) {
	// PROCESS ALL THE COMMANDS LOADED BY INIT
	slashCmds, err := s.Discord.ApplicationCommandBulkOverwrite(s.Conf.AppID, s.Conf.GuildID, slash)

	if err != nil {
		log.Println("  Command Error", err)
		// HANDLE THE ERROR
	}

	return slashCmds
}

// THIS COULD ALSO BE BROKEN OUT INTO SEPARATE FILES BUT IM MOVING EVERYTHING TO SLASH COMMANDS
func messageCreate(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// PREVENT BOT RESPONDING TO BOTS (INCLUDING ITSELF)
	if message.Author.Bot == true {
		return
	}

	// RESPOND TO USER MESSAGE IF IT CONTAINS `!HELP` OR `!BYE`
	switch {
		case strings.Contains(message.Content, "!help"):
			discord.ChannelMessageSend(message.ChannelID, "Hello World 😃")

		case strings.Contains(message.Content, "!bye"):
			discord.ChannelMessageSendReply(message.ChannelID, "Good Bye 👋", message.MessageReference)
	}
}
