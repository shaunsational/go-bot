package session

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	utils	"github.com/shaunsational/go-bot/lib/utils"
)

var (
	Conf	*Config
	Discord	*discordgo.Session
)

func Get() (*Config, *discordgo.Session) {
	return Conf, Discord
}

func init() {
	// loads values from your .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("  ERROR!!! - No .env file found")
	}
	Conf = EnvConfig()

	// create a session
	d, err := discordgo.New("Bot " + Conf.Token)
	utils.CheckNilErr(err)

	// Add the intents you will need here. the basics are covered. you can add more, pipe separated.
	// d.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentGuildMembers
	Discord = d

	d.AddHandler(ready)
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
		/*	Set the bot status.
		/	UpdateGameStatus(0, "Playing ...")
		/	UpdateWatchStatus(0, "Watching ...")
		/	UpdateListeningStatus("Listening to ...") // does NOT require the idle int.	*/

		s.UpdateListeningStatus("your commands.")
	}