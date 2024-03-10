<p align="center">
  <img alt="GO BOT" src="https://iili.io/JVyhMTx.png" />
</p>

# **Go-bot** is a base starting point for a discord bot written in go.
I've tried to break it apart into easily managed chunks of code as I'm learning [go](https://go.dev) with this project.

## Getting Started
I'm assuming you have go installed already, if not there are many better tutorials out there than what I can provide.

I recommend having [developer mode](https://www.howtogeek.com/714348/how-to-enable-or-disable-developer-mode-on-discord/)
turned on in your discord client.  After enabling this you can right click on your profile and "Copy User ID" as well as right clicking on your server you'll
be testing your bot in to "Copy Server ID" once you have these edit your .env file and paste the values on the respective lines.  Servers are also called guilds.
> AUTHOR_ID  = Your User ID that was copied
>
> GUILD_ID  = Your Server ID that was copied
>
> CLEANUP  = This will remove the commands you add every time you shutdown the bot

If CLEANUP is set to false you might have unexpected behavior if you're constantly changing the commands during development

Now you'll need to create a bot on the [Discord Developer Portal](https://discord.com/developers/applications), once that is done you'll need to populate the 
.env file with the values for your bot.  On the `General Information` tab for your bot you will find the Application ID and Public Key.  Copy those into 
your .env file so that it looks like this
> APP_ID      = 111111111111111111111
> 
> PUBLIC_KEY  = yourKEYhere
>

After that, head to the `Bot` tab and get your token, set your intents and permissions.  Copy these values in as well
> TOKEN is usually hidden and you'll need to copy it.  If you lose it, you'll need to generate a new one.
> 
> PERMISSIONS you'll need to check off the permissions you want to use and copy the generated integer

# ⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️
### BEFORE COMMITTING CHANGES YOU SHOULD IGNORE ".env" FROM GIT COMMITS

### &nbsp; 
## Next Steps
After you have all your .env variables set, open your CLI and navigate to your repo directory and run `go mod tidy`, after it finishes grabbing packages 
you'll be ready to `go run bot.go` and it should spin up an instance with your new bot.  I have it spitting out a URL you can copy to invite the bot to your
server.  

Once you've added it to your server, it has a single slash command `/command-name` that should respond with an ephemeral message of &nbsp;"..."&nbsp; 
It will also respond to DMs containing "!help" and "!bye".  if you want it to respond to messages in channels you'll need to give it intents.

### Intents
To interact with members & messages within a server you'll need to give your bot privileged gateway intents on the `Bot` tab of the 
[Discord Developer Portal](https://discord.com/developers/applications).  

Once enabled you'll need to tell your code to use them.  Open `/lib/session/session.go` and uncomment the line with `d.Identify.Intents` and 
set them appropriately.  You can find more information about them here: 
[DiscordGo Gateway Intents](https://github.com/bwmarrin/discordgo/wiki/FAQ#gateway-intents) and a full list of them here 
[List of Intents](https://github.com/bwmarrin/discordgo/blob/202785c50b9ed7366f7b92c75fea6c5e0b37f69e/structs.go#L2559)









