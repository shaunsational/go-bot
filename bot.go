//go:generate goversioninfo -icon=_src/Neon.ico
package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	session		"github.com/shaunsational/go-bot/lib/session"
	utils		"github.com/shaunsational/go-bot/lib/utils"

	commands	"github.com/shaunsational/go-bot/lib/commands"
)

var (
	conf		*session.Config
	discord		*discordgo.Session
	err 		error
)

func init() {
	conf, discord = session.Get()
}

func main() {
	slash := commands.Load()

	// open session
	err = discord.Open()
	if err != nil {
		fmt.Println("  Error opening connection,", err)
	} else {

		utils.StartupAnim(conf.AppID, conf.Permissions)
		sigalert := make(chan os.Signal, 1)
		signal.Notify(sigalert)
		<-sigalert

		if conf.Cleanup == "true" { // }
				for _, v := range slash {
					utils.LoadingText("Removing commands")
					err := discord.ApplicationCommandDelete(discord.State.User.ID, conf.GuildID, v.ID)
					if err != nil {
						fmt.Println("  - Cannot delete '%v' command: %v", v.Name, err)
					} else {
						fmt.Println("  Commands removed           ")
					}
				}
			}

		discord.Close() // close session, after function termination
	}
}