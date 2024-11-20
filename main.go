package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	botId      string
	bot        *discordgo.Session
	voiceState map[string]string
)

func main() {

	err := godotenv.Load(".env")
	token := os.Getenv("DISCORD_BOT_TOKEN")

	if err != nil {
		fmt.Println("Error loading .env file!!!")
		return
	}

	bot, err = discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := bot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	botId = u.ID

	// bot.AddHandler(ready)

	err = bot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	<-make(chan struct{})
}
